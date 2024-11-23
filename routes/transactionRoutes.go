package routes

import (
	"Expense-Tracker-go/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllTransactions)
	//r.GET("/latest", get4Transaction)
	//r.GET("/one/:id", getTransactionId)
	r.POST("/", controllers.AddTransaction)
	//r.PUT("/:id", updateTransaction)
	//r.GET("/sumIncome", getSumOfIncomeTransactionsByUserId)
	//r.GET("/sumExpense", getSumOfExpenseTransactionsByUserId)
	//r.DELETE("/:id", deleteTransactions)
}
