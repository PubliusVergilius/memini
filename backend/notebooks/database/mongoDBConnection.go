package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Open new connection
func SetupMongoDB() (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// !TODO: get URI from .env file

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	collection := client.Database("notebook-test").Collection("Notes")
	collection.Database().CreateCollection(ctx, "Notes")

	return collection, client, ctx, cancel
}

// Close the connection
func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}
