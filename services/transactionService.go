package services

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"log"
)

func AddTransaction(ctx context.Context, transaction models.Transaction) (interface{}, error) {

	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")

	result, err := collection.InsertOne(ctx, transaction)

	if err != nil {
		log.Printf("Inserting Error", err)
		return nil, err
	}

	log.Println("Inserting Success")
	return result.InsertedID, nil
}
