package repository

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindTransactionsByUser(ctx context.Context, user string) (*mongo.Cursor, error) {
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	filter := bson.M{"username": user}
	return collection.Find(ctx, filter)
}

func AddTransaction(ctx context.Context, transaction models.Transaction) (*mongo.InsertOneResult, error) {
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	return collection.InsertOne(ctx, transaction)

}
