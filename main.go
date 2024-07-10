package main

import (
	"Fuse/faq"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repoRouter := mux.NewRouter()
	log.Println("started")
	repoRouter.HandleFunc("/faq", faq.FetchFaqDetails).Methods("GET")
	http.ListenAndServe(":8020", repoRouter)
}

// var collection *mongo.Collection

// func initMongo() {
// 	var err error
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to mongo")
// 	collection = client.Database("test").Collection("FaqDetails")
// }

// func main() {
// 	// Initialize MongoDB
// 	initMongo()

// 	// Create a new context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	repoRouter := mux.NewRouter()
// 	log.Println("Server started")

// 	repoRouter.HandleFunc("/hello", FAQDetails).Methods("GET")

// 	go func() {
// 		log.Fatal(http.ListenAndServe(":8020", repoRouter))
// 	}()

// 	newFaqDetails := FaqDetails{
// 		Question:    "Where is this?",
// 		Description: "This is there",
// 		Link:        "www.com",
// 	}

// 	insertResult, err := collection.InsertOne(ctx, newFaqDetails)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted document: ", insertResult.InsertedID)

// 	var result FaqDetails
// 	filter := bson.D{{"question", "Where is this?"}}
// 	err = collection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Found document: %+v\n", result)
// }
