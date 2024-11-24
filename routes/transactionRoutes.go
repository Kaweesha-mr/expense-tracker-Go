package routes

import (
	"Expense-Tracker-go/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllTransactions)
	r.GET("/latest", controllers.Get4Transaction)
	r.GET("/:id", controllers.GetTransactionId)
	r.POST("/", controllers.AddTransaction)
	r.PUT("/:id", controllers.UpdateTransaction)
	//r.GET("/sumIncome", getSumOfIncomeTransactionsByUserId)
	//r.GET("/sumExpense", getSumOfExpenseTransactionsByUserId)
	r.DELETE("/:id", controllers.DeleteTransaction)
}
