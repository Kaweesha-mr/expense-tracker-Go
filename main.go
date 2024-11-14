package main

import (
	"Expense-Tracker-go/routes"
	"Expense-Tracker-go/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return
	}

	router := gin.Default()

	//route registering
	api := router.Group("/api/v1")
	{
		routes.UserRoutes(api)
	}

	// Step 4: Start the server
	port := os.Getenv("PORT") // Get the port from environment variable or set a default
	if port == "" {
		port = "8080" // Default port if not set in environment
	}
	log.Printf("Starting server on port %s", port)

	// Step 5: Run the Gin server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
