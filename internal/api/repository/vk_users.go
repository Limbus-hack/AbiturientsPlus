package repository

import (
	"encoding/json"
	"github.com/code7unner/vk-scrapper/config"
	"github.com/code7unner/vk-scrapper/internal/api/model"
	"github.com/code7unner/vk-scrapper/internal/api/utils"
	"net/http"
)

type VkUsers interface {
	GetVkUsers(city int) (model.VkUsers, error)
}

type vkUsersImpl struct {
	conf *config.CommonEnvConfigs
}

func NewVkUserImpl(conf config.CommonEnvConfigs) VkUsers {
	return &vkUsersImpl{&conf}
}

func (p vkUsersImpl) GetVkUsers(city int) (model.VkUsers, error) {
	userRequest := utils.UserRequest{
		Token:    p.conf.VKClientToken,
		Endpoint: "https://api.vk.com/method/users.search?sort=",
		Count:    1000,
		City:     city,
		Sort:     1,
		AgeFrom:  16,
		AgeTo:    18,
	}

	userResp, err := http.Get(userRequest.String())
	if err != nil {
		return model.VkUsers{}, err
	}
	defer userResp.Body.Close()
	var data model.VkUsers
	json.NewDecoder(userResp.Body).Decode(&data)
	return data, nil
}
