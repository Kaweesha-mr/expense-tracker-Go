package repository

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindTransactionsByUser(ctx context.Context, user string) (*mongo.Cursor, error) {
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	filter := bson.M{"username": user}
	return collection.Find(ctx, filter)
}

func FindLatestTransactionsByUser(ctx context.Context, user string, opt *options.FindOptions) (*mongo.Cursor, error) {
	format := bson.M{"username": user}
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	return collection.Find(ctx, format, opt)
}

func AddTransaction(ctx context.Context, transaction models.Transaction) (*mongo.InsertOneResult, error) {
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	return collection.InsertOne(ctx, transaction)

}

func FindTransactionById(ctx context.Context, id primitive.ObjectID) (*mongo.SingleResult, error) {
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	format := bson.M{"_id": id}
	return collection.FindOne(ctx, format), nil
}

func UpdateTransactionById(ctx context.Context, id primitive.ObjectID, data bson.M) (*mongo.UpdateResult, error) {
	// Step 1: Define the filter and update query
	filter := bson.M{"_id": id}    // Filter to find the document by ID
	update := bson.M{"$set": data} // Use $set to update only the specified fields

	// Step 2: Execute the UpdateOne query
	collection := utils.Client.Database("ExpenseTracker").Collection("Transaction")
	result, err := collection.UpdateOne(ctx, filter, update)

	// Step 3: Handle errors
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	// Step 4: Return the UpdateResult
	return result, nil
}
