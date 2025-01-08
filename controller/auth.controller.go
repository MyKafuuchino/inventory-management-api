package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/config"
	"inventory-management/entity"
	"inventory-management/service"
	"inventory-management/utils"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user service.Login

	appConfig := config.GlobalAppConfig
	jwtService := utils.NewJwtService([]byte(appConfig.SecretKey))

	if err := ctx.ShouldBindJSON(&user); err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusBadRequest, err.Error()))
		return
	}

	loggedUser, err := c.authService.Login(&user)

	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}

	token, err := jwtService.GenJwtToken(loggedUser.ID)
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, entity.NewResponseSuccess("Success login", gin.H{
		"username":  loggedUser.Username,
		"full_name": loggedUser.FullName,
		"token":     token,
	}))
}
