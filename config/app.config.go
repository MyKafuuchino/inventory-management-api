package config

import "inventory-management/utils"

type AppConfig struct {
	AppPort string
}

var GlobalAppConfig AppConfig

func InitAppConfig() {
	GlobalAppConfig = AppConfig{
		AppPort: utils.GetEnv("APP_PORT", "8080"),
	}
}
