package mongo_client

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	var err error
	Client, err = GetMongoClient()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func GetMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func CloseMongoClient() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}
	}
}
