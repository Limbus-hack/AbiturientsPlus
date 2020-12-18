package config

import (
	"github.com/caarlos0/env"
	"log"
)

type CommonEnvConfigs struct {
	// Logging level
	LogLevel string `json:"LOG_LEVEL" env:"LOG_LEVEL" envDefault:"debug"`

	// Service configs
	ServiceName string `json:"SERVICE_NAME" env:"SERVICE_NAME" envDefault:"rest-api-template"`

	// Server configs
	ServerPort string `json:"SERVER_PORT" env:"SERVER_PORT" envDefault:"8082"`

	// PostgreSQL configs
	PostgresURL      string `json:"POSTGRES_URL" env:"POSTGRES_URL" envDefault:"localhost"`
	PostgresPort     int    `json:"POSTGRES_PORT" env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser     string `json:"POSTGRES_USER" env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPassword string `json:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDB       string `json:"POSTGRES_DB" env:"POSTGRES_DB" envDefault:"sites"`

	// VK configs
	VkAppID      int    `json:"VK_APP_ID" env:"VK_APP_ID"`
	VkPrivateKey string `json:"VK_PRIVATE_KEY" env:"VK_PRIVATE_KEY"`
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
