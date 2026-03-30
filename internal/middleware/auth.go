package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminAuth creates a middleware that validates admin authentication
func AdminAuth(adminSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if adminSecret == "" {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Admin functionality is disabled"})
			c.Abort()
			return
		}

		token := c.GetHeader("Authorization")
		if token != adminSecret {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Next()
	}
}
