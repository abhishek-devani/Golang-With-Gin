package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type manager struct {
	Connection *mongo.Client
	Ctx context.Context
	Cancel context.CancelFunc
}

func dbConnection() {

	// Set your MongoDB connection string
	connectionURI := "mongodb://localhost:27017"

	//  Set up client options
	clientOptions := options.Client().ApplyURI(connectionURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	// Connect to Mongo
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = manager {
		Connection: client,
		Ctx: ctx,
		Cancel: cancel,
	}

	fmt.Println("Connected...")

}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func main() {

	dbConnection()
}
