package middleware

import (
	"Expense-Tracker-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type contextKey string

const UserClaimsKey = contextKey("userName")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Get the authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort() // Stop further execution
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		info, err := utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Expired token"})
			c.Abort()
			return
		}

		//add claims to the request context
		c.Set(string(UserClaimsKey), info.Username)

		c.Next()

	}

}
