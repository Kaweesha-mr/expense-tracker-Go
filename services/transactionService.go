package services

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/repository"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func AddTransaction(ctx context.Context, transaction models.Transaction) (interface{}, error) {

	result, err := repository.AddTransaction(ctx, transaction)

	if err != nil {
		log.Printf("Inserting Error", err)
		return nil, err
	}

	log.Println("Inserting Success")
	return result.InsertedID, nil
}

func GetAllTransactions(ctx context.Context, user string) ([]models.Transaction, error) {

	cursor, err := repository.FindTransactionsByUser(ctx, user)
	if err != nil {

		logrus.WithFields(logrus.Fields{
			"user":  user,
			"error": err,
		}).Error("Failed with transactions")

		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

			logrus.WithError(err).Error("Failed to close the mongoDB Cursor")

		}
	}(cursor, ctx)

	var transactions []models.Transaction

	for cursor.Next(ctx) {

		//context error handling
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		var transaction models.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			logrus.WithFields(logrus.Fields{
				"user":  user,
				"error": err,
			}).Error("Failed to decode transaction")
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  user,
			"error": err,
		}).Error("Error during cursor iteration")
		return nil, fmt.Errorf("error during cursor iteration: %w", err)
	}

	// Handle no results
	if len(transactions) == 0 {
		logrus.WithField("user", user).Info("No transactions found for user")
		return []models.Transaction{}, nil
	}

	logrus.WithFields(logrus.Fields{
		"user":             user,
		"transactionCount": len(transactions),
	}).Info("Successfully fetched transactions")

	return transactions, nil
}
