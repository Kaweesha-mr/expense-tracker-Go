package repository

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = utils.Client.Database("ExpenseTracker").Collection("Transaction")

func FindTransactionsByUser(ctx context.Context, user string) (*mongo.Cursor, error) {
	filter := bson.M{"username": user}
	return collection.Find(ctx, filter)
}

func AddTransaction(ctx context.Context, transaction models.Transaction) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(ctx, transaction)

}
