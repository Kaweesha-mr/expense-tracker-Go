package controllers

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTransaction(c *gin.Context) {

	ctx := c.Request.Context()

	User, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User Claims Not Found"})
		return
	}

	//User, ok := User.(string)
	//if !ok {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid or missing user claims"})
	//	return
	//}

	var Transaction models.Transaction
	Transaction.UserName = User.(string)

	if err := c.ShouldBind(&Transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Binding data Error", "Details": err.Error()})
		return
	}

	result, err := services.AddTransaction(ctx, Transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB Error", "Details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Object Id": result})

}
