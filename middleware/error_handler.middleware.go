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
			var e *utils.CustomError
			switch {
			case errors.As(err.Err, &e):
				ctx.JSON(e.StatusCode, utils.NewResponseError(e.Message, e.Errors...))
			default:
				ctx.JSON(http.StatusInternalServerError, utils.NewResponseError(http.StatusText(http.StatusInternalServerError)))
			}
		}
	}
}
