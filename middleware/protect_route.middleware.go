package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"inventory-management/config"
	"inventory-management/database"
	"inventory-management/entity"
	"inventory-management/utils"
	"net/http"
	"strings"
)

func ProtectRoute(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretKey := config.GlobalAppConfig.SecretKey
		token := ctx.GetHeader("Authorization")
		if token == "" {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Authorization header is empty"))
			ctx.Abort()
			return
		}

		tokenPart := strings.Split(token, " ")
		if len(tokenPart) != 2 || tokenPart[0] != "Bearer" {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid token format"))
			ctx.Abort()
			return
		}

		parsedToken, err := jwt.Parse(tokenPart[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !parsedToken.Valid {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid token"))
			ctx.Abort()
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid token claims"))
			ctx.Abort()
			return
		}

		fmt.Println("Claims ", claims)

		userID, ok := claims["userId"].(float64)
		if !ok {
			err = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "UserId not found in token"))
			ctx.Abort()
			return
		}

		user := &entity.User{}
		result := database.DB.Table("users").Select("id", "username", "role").Where("id = ?", userID).First(user)

		if result.Error != nil {
			_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "User not found"))
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if role == user.Role {
				ctx.Set("user", user)
				ctx.Next()
			} else {
				_ = ctx.Error(utils.NewCustomError(http.StatusUnauthorized, "Invalid role"))
				ctx.Abort()
				return
			}
		}
	}
}
