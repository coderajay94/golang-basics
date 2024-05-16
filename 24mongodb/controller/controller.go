package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName             string = "netflix"
	collectionName     string = "watchlist"
	dbConnectionString string = "mongodb+srv://ajakumar:admin@cluster0.hecejb8.mongodb.net/"
)

var collection *mongo.Collection

//connect with mongo db

// only runs one time, only when application is starting
func init() {

	//client option
	clientOptions := options.Client().ApplyURI(dbConnectionString)

	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("error occured while connecting to mongodb database")
		log.Fatal(err)
	}

	fmt.Println("client connected successfully")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("collection instance is ready")
}
