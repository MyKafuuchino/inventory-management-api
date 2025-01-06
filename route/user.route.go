package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func UserRoute(ctx *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	user := ctx.Group("/users")
	{
		user.GET("/", userController.GetAllUsers)
		user.GET("/:id", userController.GetUserById)
		user.POST("/", userController.CreateNewUser)
	}
}
