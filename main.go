package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joalvarezdev/go-gpt/database"
	"github.com/joalvarezdev/go-gpt/handlers"
)

func main() {

	database.Init()

	router := mux.NewRouter()

	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	log.Println("Server running on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}