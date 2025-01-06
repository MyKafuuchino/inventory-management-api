package config

import "inventory-management/utils"

type DbConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

var GlobalDbConfig DbConfig

func InitDbConfig() {
	GlobalDbConfig = DbConfig{
		DbHost:     utils.GetEnv("DB_HOST", "127.0.0.1"),
		DbPort:     utils.GetEnv("DB_PORT", "3306"),
		DbUser:     utils.GetEnv("DB_USER", "root"),
		DbPassword: utils.GetEnv("DB_PASSWORD", ""),
		DbName:     utils.GetEnv("DB_NAME", "inventory-management"),
	}
}
