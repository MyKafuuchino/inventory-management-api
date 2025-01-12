package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/config"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func AuthRoute(ctx *gin.RouterGroup) {
	appConfig := config.GlobalAppConfig

	authRepository := repository.NewAuthRepository(database.DB)
	authService := service.NewAuthService(authRepository, []byte(appConfig.SecretKey))
	authController := controller.NewAuthController(authService)

	auth := ctx.Group("/auth")
	{
		auth.POST("/login", authController.Login)
	}
}
