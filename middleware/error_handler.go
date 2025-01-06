package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inventory-management/entity"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err != nil {
			var e *entity.CustomError
			switch {
			case errors.As(err.Err, &e):
				ctx.JSON(e.StatusCode, entity.NewResponseError(e.Message))
			default:
				ctx.JSON(http.StatusInternalServerError, entity.NewResponseError(http.StatusText(http.StatusInternalServerError)))
			}
		}
	}
}
