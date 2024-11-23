package controllers

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func AddTransaction(c *gin.Context) {

	ctx := c.Request.Context()

	User, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User Claims Not Found"})
		return
	}

	var Transaction models.Transaction
	Transaction.UserName = User.(string)

	if err := c.ShouldBind(&Transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Binding data Error", "Details": err.Error()})
		return
	}

	//validating User Input

	validate := validator.New()
	if err := validate.Struct(Transaction); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		errorMessages := make(map[string]string)

		for _, validationError := range validationErrors {
			field := validationError.Field() // Field name
			tag := validationError.Tag()     // Validation rule
			errorMessages[field] = "Validation failed for " + field + " (" + tag + ")"
		}

		c.JSON(http.StatusBadRequest, gin.H{"ValidationErrors": errorMessages})
		return
	}

	result, err := services.AddTransaction(ctx, Transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB Error", "Details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Object Id": result})

}
