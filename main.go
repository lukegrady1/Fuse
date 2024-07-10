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
