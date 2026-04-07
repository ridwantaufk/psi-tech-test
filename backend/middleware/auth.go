package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridwantaufk/psi-tech-test/utils"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")

		if err != nil {
			tokenStr = c.GetHeader("Authorization")
			if tokenStr == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}
		}
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
