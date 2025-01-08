package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func AuthRoute(ctx *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(database.DB)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	auth := ctx.Group("/auth")
	{
		auth.POST("/login", authController.Login)
	}
}
