package service

import (
	"encoding/json"
	"fmt"
	"github.com/code7unner/vk-scrapper/config"
	"github.com/code7unner/vk-scrapper/internal/api/model"
	"github.com/code7unner/vk-scrapper/internal/api/utils"
	"net/http"
	"regexp"
	"strings"
)

func GetVkUsers(city int, conf *config.CommonEnvConfigs) ([]model.VkUsers, error) {
	userRequest := utils.UserRequest{
		Token:    conf.VKClientToken,
		Endpoint: "https://api.vk.com/method/users.search?sort=",
		Count:    1000,
		City:     city,
		Sort:     1,
		AgeFrom:  16,
		AgeTo:    18,
		Fields:   "sex,interests",
	}
	var dataSlice []model.VkUsers
	for i := 0; i < 33; i++ {
		userResp, err := http.Get(userRequest.String(utils.ToChar(i)))
		if err != nil {
			return nil, err
		}
		var data model.VkUsers
		json.NewDecoder(userResp.Body).Decode(&data)
		dataSlice = append(dataSlice, data)
		userResp.Body.Close()
	}
	return dataSlice, nil
}

func GetVkUserSubs(userId int64, conf *config.CommonEnvConfigs) (model.VkUserSubscriptions, error) {
	subsRequest := utils.SubsRequest{
		Token:    conf.VKClientToken,
		Endpoint: "https://api.vk.com/method/users.getSubscriptions?extended=",
		Count:    200,
		UserId:   userId,
		Extended: 1,
	}
	test := subsRequest.String()
	print(test)
	subsResp, err := http.Get(subsRequest.String())
	if err != nil {
		return model.VkUserSubscriptions{}, err
	}
	var data model.VkUserSubscriptions
	json.NewDecoder(subsResp.Body).Decode(&data)
	return data, nil
}

func BulkGetVkUserSubs(users []model.VkUsers, conf *config.CommonEnvConfigs) (userStrings []string, err error) {
	reg, err := regexp.Compile("[^a-zA-Zа-яА-Я0-9]+")
	if err != nil {
		return userStrings, err
	}
	for _, user := range users {
		for _, item := range user.Response.Items {

			userSubs, err := GetVkUserSubs(item.Id, conf)
			if err != nil {
				return nil, err
			}
			groupsName := make([]string, 0, len(userSubs.SubscriptionsResponse.SubscriptionItems))
			for _, sub := range userSubs.SubscriptionsResponse.SubscriptionItems {
				s := strings.ToLower(sub.Name)
				s = reg.ReplaceAllString(s, "")
				s = strings.ReplaceAll(s, " ", "")
				groupsName = append(groupsName, s)
			}
			interests := strings.Split(item.Interests, ",")
			interestsNames := make([]string, 0, len(interests))
			for _, interest := range interests {
				i := strings.ToLower(interest)
				i = reg.ReplaceAllString(i, "")
				i = strings.ReplaceAll(i, " ", "")
				interestsNames = append(interestsNames, i)
			}

			sex := ""
			if item.Sex == 1 {
				sex = "female"
			} else {
				sex = "male"
			}
			if len(groupsName) > 0 || interestsNames[0] != "" {
				userStrings = append(userStrings, fmt.Sprintf("| %s %s %s",
					strings.Join(groupsName, " "),
					strings.Join(interestsNames, " "),
					sex))
			}
		}
	}
	return userStrings, nil
}
