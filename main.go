package main

import (
	"Expense-Tracker-go/utils"
	"log"
)

func main() {

	err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return
	}

}
