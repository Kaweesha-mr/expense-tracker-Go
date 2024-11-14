package utils

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

}

func ConnectDB() error {
	uri := os.Getenv("MONGODB_URI")
	var err error

	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if uri == "" {
		return fmt.Errorf("failed to connect to Mongodb %v", err)

	}

	//after connection setted up this is to test that connection
	if err = Client.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("Mongodb Connection Failed: %v", err)
	}

	log.Printf("Successfully connected to Mongodb")
	return nil
}
