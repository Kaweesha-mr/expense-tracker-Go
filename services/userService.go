package services

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetUserByUsername(username string) (*models.User, error) {

	collection := utils.Client.Database("ExpenseTracker").Collection("Users")

	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		log.Println("Error finding user:", err)
		return nil, err
	}

	return &user, nil

}
