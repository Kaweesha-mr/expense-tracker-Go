package services

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/repository"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func Get4Transactions(c *gin.Context, user string) ([]models.Transaction, error) {
	// Prepare query options
	opt := options.Find()
	opt.SetLimit(4)

	// Fetch transactions using the repository
	cursor, err := repository.FindLatestTransactionsByUser(c.Request.Context(), user, opt)
	if err != nil {
		logrus.WithField("user", user).WithError(err).Error("Failed to fetch transactions")
		return nil, err
	}

	// Ensure the cursor is closed
	defer func() {
		if err := cursor.Close(c.Request.Context()); err != nil {
			logrus.WithError(err).Error("Failed to close the MongoDB cursor")
		}
	}()

	var transactions []models.Transaction

	// Decode transactions
	for cursor.Next(c.Request.Context()) {
		var transaction models.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			logrus.WithField("user", user).WithError(err).Error("Failed to decode transaction")
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	// Handle errors during cursor iteration
	if err := cursor.Err(); err != nil {
		logrus.WithField("user", user).WithError(err).Error("Error during cursor iteration")
		return nil, err
	}

	// Handle no results
	if len(transactions) == 0 {
		logrus.WithField("user", user).Info("No transactions found for user")
		return transactions, nil
	}

	// Log success and return transactions
	logrus.WithFields(logrus.Fields{
		"user":             user,
		"transactionCount": len(transactions),
	}).Info("Successfully fetched transactions")

	return transactions, err
}
