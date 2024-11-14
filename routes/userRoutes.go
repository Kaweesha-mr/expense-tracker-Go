package routes

import (
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.RouterGroup) {
	r.POST("/login", login)
	r.POST("/signup", signup)

}
