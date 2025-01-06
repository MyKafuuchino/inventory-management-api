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
	if len(users) == 0 {
		err = ctx.Error(entity.NewCustomError(http.StatusNotFound, "No users found create new user"))
		return
	}
	ctx.JSON(http.StatusOK, entity.NewResponseSuccess[[]entity.User]("Success get user", users))
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

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusBadRequest, err.Error()))
		return
	}
	createdUser, err := c.userService.CreateNewUser(&user)
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, entity.NewResponseSuccess[*entity.User]("Success create user", createdUser))
}
