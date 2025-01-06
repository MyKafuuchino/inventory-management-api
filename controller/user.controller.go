package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/entity"
	"inventory-management/service"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := c.userService.GetUserById(userId)
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, entity.NewResponseSuccess[entity.User]("Success get user", user))
}
