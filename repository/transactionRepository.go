package repository

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
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
