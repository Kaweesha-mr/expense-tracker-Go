package services

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func GetAllTransactions(ctx context.Context, user string) ([]models.Transaction, error) {
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")

	filter := bson.M{"username": user}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {

		log.Println("Finding Error:", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var transactions []models.Transaction

	for cursor.Next(ctx) {
		var transaction models.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
