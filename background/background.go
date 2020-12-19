package background

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/code7unner/vk-scrapper/background/models"
	"github.com/code7unner/vk-scrapper/internal/api/model"
	"github.com/code7unner/vk-scrapper/internal/app"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	bufSize = 100000
)

type Background struct {
	wg  *sync.WaitGroup
	app *app.App
}

func New(app *app.App) *Background {
	return &Background{
		wg:  &sync.WaitGroup{},
		app: app,
	}
}

func (b *Background) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Stop()
			//"inf_bu", "ege_matn", "physics_100", "ege", "egeoge_math"
			groupIDs := []string{"inf_bu", "ege_matn", "physics_100", "ege", "egeoge_math"}
			vkID := make(chan int, bufSize)
			wg := sync.WaitGroup{}
			for _, id := range groupIDs {
				offset := 0
				wg.Add(1)
				go b.getVkIDs(vkID, id, offset, &wg)
				go func() {
					wg.Wait()
					close(vkID)
				}()
			}

			b.wg.Add(1)
			go b.sender(vkID)
			b.wg.Wait()
		}
	}
}

func (b *Background) sender(vkID chan int) {
	for {
		select {
		case id, more := <-vkID:
			if !more {
				b.wg.Done()
				return
			}

			profile, err := b.getVkProfile(id)
			if err != nil {
				continue
			}
			group, err := b.getVkGroup(id)
			if err != nil {
				continue
			}

			str, err := b.buildString(profile, group)
			if err != nil {
				continue
			}

			res, err := b.app.Vws.Predict(str)
			if err != nil {
				continue
			}

			user := &model.User{
				ID:         profile.Response[0].ID,
				Name:       profile.Response[0].FirstName,
				LastName:   profile.Response[0].LastName,
				Region:     profile.Response[0].City.ID,
				Prediction: res,
				Status:     "new",
			}

			_, err = b.app.Repo.Users.Create(b.app.Ctx, user)
			if err != nil {
				b.app.Log.Warn(err)
				continue
			}
		}
	}
}

func (b *Background) buildString(profile *models.VkUserModel, group *models.VkGroupModel) (string, error) {
	reg, err := regexp.Compile("[^a-zа-я0-9]+")
	if err != nil {
		return "", err
	}

	interests := strings.Split(profile.Response[0].Interests, ",")
	intesertsStrings := make([]string, 0, len(interests))
	for _, i := range interests {
		i = strings.ToLower(i)
		i = reg.ReplaceAllString(i, "")
		i = strings.ReplaceAll(i, " ", "")
		intesertsStrings = append(intesertsStrings, i)
	}

	groups := group.Response.Items
	groupsStrings := make([]string, 0, len(groups))
	for _, g := range groups {
		gn := strings.ToLower(g.Name)
		gn = reg.ReplaceAllString(gn, "")
		gn = strings.ReplaceAll(gn, " ", "")
		groupsStrings = append(groupsStrings, gn)
	}

	var Sex string
	if profile.Response[0].Sex == 1 {
		Sex = "female"
	} else {
		Sex = "male"
	}

	formatedString := fmt.Sprintf("| %s %s %s\n",
		strings.Join(groupsStrings, " "),
		strings.Join(intesertsStrings, ""),
		Sex)

	return formatedString, nil
}

func (b *Background) getVkGroup(id int) (*models.VkGroupModel, error) {
	url := "https://api.vk.com/method/users.getSubscriptions?user_id=%d&extended=1&count=200&v=5.126&access_token=%s"

	url = fmt.Sprintf(url, id, b.app.Conf.VkServiceToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var data models.VkGroupModel
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	resp.Body.Close()

	return &data, nil
}

func (b *Background) getVkProfile(id int) (*models.VkUserModel, error) {
	url := "https://api.vk.com/method/users.get?user_id=%d&fields=verified,interests,sex,city&extended=1&count=1000&v=5.126&access_token=%s"

	url = fmt.Sprintf(url, id, b.app.Conf.VkServiceToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var data models.VkUserModel
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	resp.Body.Close()

	if data.Response == nil {
		return nil, errors.New("invalid user")
	}

	if data.Response[0].IsClosed || data.Response[0].Deactivated != "" {
		return nil, errors.New("invalid user")
	}

	return &data, nil
}

func (b *Background) getVkIDs(vkID chan int, id string, offset int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		url := "https://api.vk.com/method/groups.getMembers?group_id=%s&sort=id_asc&offset=%d&count=1000&v=5.126&access_token=%s"

		url = fmt.Sprintf(url, id, offset, b.app.Conf.VkServiceToken)
		resp, err := http.Get(url)
		if err != nil {
			break
		}

		offset += 1000

		var data models.VkIDModel
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			break
		}
		resp.Body.Close()

		if len(data.Response.Items) == 0 {
			break
		}

		for _, item := range data.Response.Items {
			vkID <- item
		}
	}
}
