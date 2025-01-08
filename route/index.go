package route

import "github.com/gin-gonic/gin"

func InitRoute(ctx *gin.Engine) {
	api := ctx.Group("/api")
	UserRoute(api)
	AuthRoute(api)
}
