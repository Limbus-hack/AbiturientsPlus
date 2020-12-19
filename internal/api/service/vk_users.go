package service

import (
	"encoding/json"
	"github.com/code7unner/vk-scrapper/config"
	"github.com/code7unner/vk-scrapper/internal/api/model"
	"github.com/code7unner/vk-scrapper/internal/api/utils"
	"net/http"
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
