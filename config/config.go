package config

import (
	"github.com/caarlos0/env"
	"log"
)

type CommonEnvConfigs struct {
	// Logging level
	LogLevel string `json:"LOG_LEVEL" env:"LOG_LEVEL" envDefault:"debug"`

	// Service configs
	ServiceName string `json:"SERVICE_NAME" env:"SERVICE_NAME" envDefault:"vk-scrapper"`

	// Server configs
	ServerPort string `json:"SERVER_PORT" env:"SERVER_PORT" envDefault:"8082"`

	// PostgreSQL configs
	PostgresDBStr    string `json:"POSTGRES_DB_STR" env:"POSTGRES_DB_STR" envDefault:"postgresql://postgres:postgres@128.199.77.142:5432/vk_users"`

	// VK configs
	VkAppID        int    `json:"VK_APP_ID" env:"VK_APP_ID"`
	VkPrivateKey   string `json:"VK_PRIVATE_KEY" env:"VK_PRIVATE_KEY"`
	VkServiceToken string `json:"VK_SERVICE_TOKEN" env:"VK_SERVICE_TOKEN"`
	VKClientToken  string `json:"VK_CLIENT_TOKEN" env:"VK_CLIENT_TOKEN"`
}

func GetCommonEnvConfigs() CommonEnvConfigs {
	envConfig := CommonEnvConfigs{}
	err := env.Parse(&envConfig)
	if err != nil {
		log.Panicf("Error parse env config:%s", err)
		return envConfig
	}
	return envConfig
}
