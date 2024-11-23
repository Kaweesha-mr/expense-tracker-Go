package utils

import "github.com/gin-gonic/gin"

// RespondWithError Helper: Error Response
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
	c.Abort()
}

// RespondWithJSON Helper: JSON Response
func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, gin.H{"data": payload})
}
