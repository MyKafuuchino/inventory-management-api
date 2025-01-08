package config

import "inventory-management/utils"

type AppConfig struct {
	AppPort   string
	SecretKey string
}

var GlobalAppConfig AppConfig

func InitAppConfig() {
	GlobalAppConfig = AppConfig{
		AppPort:   utils.GetEnv("APP_PORT", "8080"),
		SecretKey: utils.GetEnv("JWT_SECRET_KEY", ""),
	}
}
