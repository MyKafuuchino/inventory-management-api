package main

import (
	"github.com/gin-gonic/gin"
	"inventory-management/config"
	"inventory-management/database"
	"inventory-management/database/seeder"
	"inventory-management/entity"
	"inventory-management/middleware"
	"inventory-management/route"
	"net/http"
)

func main() {
	config.InitEnvConfig()

	database.InitDatabase()
	seeder.SeedUser()

	appConfig := config.GlobalAppConfig

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, entity.NewResponseError("Not Found"))
	})

	route.InitRoute(r)

	err := r.Run(appConfig.AppPort)
	if err != nil {
		panic("Fail to start gin server: " + err.Error())
	}
}
