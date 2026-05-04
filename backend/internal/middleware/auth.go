package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"hobby-blog/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			c.Abort()
			return
		}

		headerParts := strings.SplitN(authorizationHeader, " ", 2)

		if len(headerParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			c.Abort()
			return
		}

		bearerType := headerParts[0]
		tokenString := headerParts[1]

		if bearerType != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token type",
			})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
