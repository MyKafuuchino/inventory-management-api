package config

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	InitAppConfig()
	InitDbConfig()
}
