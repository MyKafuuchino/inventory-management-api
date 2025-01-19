package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"inventory-management/config"
	"inventory-management/database"
	"inventory-management/entity"
	"net/http"
	"strings"
	"time"
)

func ProtectRoute(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretKey := config.GlobalAppConfig.SecretKey
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			ctx.Abort()
			return
		}

		tokenPart := strings.Split(token, " ")
		if len(tokenPart) != 2 || tokenPart[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
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
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				ctx.Abort()
				return
			}
		}

		userID, ok := claims["userId"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "UserId not found in token"})
			ctx.Abort()
			return
		}

		user := &entity.User{}
		result := database.DB.Table("users").Select("id", "username", "role").Where("id = ?", userID).First(user)
		if result.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			ctx.Abort()
			return
		}

		isAuthorized := false
		for _, role := range roles {
			if role == user.Role {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient role permissions"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
