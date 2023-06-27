package main

import (
	"context"
	"github.com/mongodb/mongo-driver/mongo"
	"github.com/mongodb/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitClient() *mongo.Client {
	// Replace this with your actual MongoDB connection string
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://yourmongodburl"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
