package middleware

import (
	"net/http"
	"os"
	"strings"

	"multishop/config"
	"multishop/models"
	"multishop/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) == 2 && parts[0] == "Bearer" {
			authHeader = parts[1]
		}

		claims := &utils.JWTClaims{}

		token, err := jwt.ParseWithClaims(
			authHeader,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		var user models.User

		if err := config.DB.
			Where("id = ? AND shop_id = ?", claims.UserID, claims.ShopID).
			First(&user).Error; err != nil {

			utils.Error(c, http.StatusUnauthorized, "user not found")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("shop_id", claims.ShopID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
