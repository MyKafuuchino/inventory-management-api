package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/service"
	"inventory-management/utils"
	"inventory-management/validation"
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
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[[]entity.User]("Success get user", users))
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := c.userService.GetUserByID(userId)

	if err != nil {
		err = ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[*entity.User]("Success get user", user))
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var userRequest *model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		err = ctx.Error(err)
		return
	}
	if err := validation.ValidationHandler(userRequest); err != nil {
		err = ctx.Error(err)
		return
	}
	createdUser, err := c.userService.CreateNewUser(userRequest)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, utils.NewResponseSuccess[*entity.User]("Success create user", createdUser))
}

func (c *UserController) DeleteUserByID(ctx *gin.Context) {
	userId := ctx.Param("id")
	err := c.userService.DeleteUserByID(userId)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[*entity.User]("Success delete user", nil))
}
