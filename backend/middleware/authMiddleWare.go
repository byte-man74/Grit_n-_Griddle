package AuthMiddleware

import (
	AuthUtils "github.com/byte-man74/Grit_n-_Griddle/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from headers
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		// Validate and extract user information from the token
		user, err := AuthUtils.ValidateAndExtractUser(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Set user information in the context
		c.Set("user", user)

		// Continue to the next middleware or route handler if there is
		c.Next()
	}
}
