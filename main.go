package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	repoRouter := mux.NewRouter()
	log.Println("started")
	repoRouter.HandleFunc("/hello", UserAuthentication).Methods("GET")
	http.ListenAndServe(":8020", repoRouter)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	// Get a handle for your database and collection
	collection := client.Database("test").Collection("faq")

}

func UserAuthentication(w http.ResponseWriter, r *http.Request) {
	log.Println("yest")
}
