package main

import (
	"github.com/gin-gonic/gin"
	"inventory-management/config"
	"inventory-management/database"
	"inventory-management/middleware"
	"inventory-management/route"
)

func main() {
	config.InitEnvConfig()

	database.InitDatabase()

	appConfig := config.GlobalAppConfig

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	route.InitRoute(r)

	err := r.Run(appConfig.AppPort)
	if err != nil {
		panic("Fail to start gin server: " + err.Error())
	}
}
