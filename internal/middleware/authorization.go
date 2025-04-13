package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware struct{}

func InitAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Check for a specific header
		if c.GetHeader("Authorization") == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Continue to the next middleware/handler
		c.Next()
	}
}
