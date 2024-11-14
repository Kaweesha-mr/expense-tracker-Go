package controllers

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/services"
	"Expense-Tracker-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid  input", "details": err.Error()})
	}

	err := services.SaveUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

	}

	// Generate JWT access and refresh tokens
	accessToken, err := utils.GenerateAccessToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "error Generating"})
	}

	// Send the tokens in the response
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
