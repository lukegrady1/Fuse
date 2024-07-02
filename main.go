package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repoRouter := mux.NewRouter()
	log.Println("started")
	repoRouter.HandleFunc("/hello", UserAuthentication).Methods("GET")
	http.ListenAndServe(":8020", repoRouter)
}

func UserAuthentication(w http.ResponseWriter, r *http.Request) {
	log.Println("yest")
}
