package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))

	if err != nil {
		panic(err)
	}
	log.Println("successfully connected to mongodb")
	var db = client.Database("db")
	return db
}
