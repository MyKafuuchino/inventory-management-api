package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inventory-management/utils"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err != nil {
			var customError *utils.CustomError
			if errors.As(err, &customError) {
				ctx.JSON(customError.StatusCode, utils.NewResponseError(customError.Message, customError.Errors...))
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
	}
}
