package services

import (
	"Expense-Tracker-go/models"
	"Expense-Tracker-go/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func SaveUser(username, password string) error {

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err

	}

	user := models.User{
		Username: username,
		Password: hashedPassword,
	}

	collection := utils.Client.Database("ExpenseTracker").Collection("Users")

	var existingUser models.User
	err = collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		// If the user already exists
		return errors.New("user already exists")
	} else if err.Error() != "mongo: no documents in result" {
		// Log unexpected errors
		log.Println("Error checking for existing user:", err)
		return err
	}

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error inserting User:", err)
		return err
	}

	log.Println("User Created Successfully")
	return nil
}
