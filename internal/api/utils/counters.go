package utils

import "github.com/code7unner/vk-scrapper/internal/api/model"

func CountUsers(users []model.VkUsers) int64 {
	var globalCount int64
	for i := 0; i < 33; i++ {
		globalCount += users[i].Response.Count
	}
	return globalCount
}
