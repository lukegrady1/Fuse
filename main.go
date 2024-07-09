package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FaqDetails struct {
	Question    string `bson:"question"`
	Description string `bson:"description"`
	Link        string `bson:"link"`
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	fmt.Println("Connected")

	// Get a handle for your database and collection
	collection := client.Database("test").Collection("FaqDetails")

	newFaqDetails := FaqDetails{
		Question:    "Where is this?",
		Description: "This is there",
		Link:        "www.com",
	}

	insertResult, err := collection.InsertOne(ctx, newFaqDetails)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert a single document: ", insertResult.InsertedID)

	var result FaqDetails
	filter := bson.D{{"question", "Where is this?"}}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

}
