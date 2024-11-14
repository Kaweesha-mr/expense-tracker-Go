package routes

import (
	"Expense-Tracker-go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignUp)

}
