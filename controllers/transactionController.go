package controllers

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/services"
	"Expense-Tracker-go/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddTransaction(c *gin.Context) {

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

	result, err := services.AddTransaction(c.Request.Context(), Transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB Error", "Details": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Object Id": result})

}

func GetAllTransactions(c *gin.Context) {

	User, exists := c.Get("userName")
	if !exists {

		utils.RespondWithError(c, http.StatusInternalServerError, "User Claims Not Found")
		return
	}

	transactions, err := services.GetAllTransactions(c.Request.Context(), User.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, transactions)

}

func Get4Transaction(c *gin.Context) {

	User, exists := c.Get("userName")
	if !exists {

		utils.RespondWithError(c, http.StatusInternalServerError, "User Claims Not Found")
		return
	}

	transactions, err := services.Get4Transactions(c, User.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, transactions)

}

func GetTransactionId(c *gin.Context) {

	objectId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid Transaction ID format")
		return
	}

	User, exists := c.Get("userName")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User Claims Not Found")
		return
	}

	transaction, err := services.GetTransactionById(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "transaction not found" {
			utils.RespondWithError(c, http.StatusNotFound, "Transaction not found")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Error fetching transaction data")
		return
	}

	// Step 5: Return the transaction data as a JSON response
	utils.RespondWithJSON(c, http.StatusOK, gin.H{
		"user":        User,
		"transaction": transaction,
	})
}
