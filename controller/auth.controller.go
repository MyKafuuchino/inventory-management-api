package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/model"
	"inventory-management/service"
	"inventory-management/utils"
	"inventory-management/validation"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginBody model.Login

	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		err = ctx.Error(err)
		return
	}

	if err := validation.ValidationHandler[*model.Login](&loginBody); err != nil {
		err = ctx.Error(err)
		return
	}

	loggedUser, token, err := c.authService.Login(&loginBody)

	if err != nil {
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success login", gin.H{
		"full_name": loggedUser.FullName,
		"username":  loggedUser.Username,
		"token":     token,
	}))
}
