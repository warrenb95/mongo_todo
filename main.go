package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/warrenb95/mongo_todo/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client for mongoDB connection
var Client *mongo.Client

func main() {
	fmt.Println("Starting...")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	Client, _ = mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB at localhost:27017")

	routes.Client = Client

	router := mux.NewRouter()
	router.HandleFunc("/todo", routes.CreateTodoEndPoint).Methods("POST")
	router.HandleFunc("/todo", routes.GetAllTodosEndPoint).Methods("GET")
	router.HandleFunc("/todo/{id}", routes.GetTodoEndpoint).Methods("GET")
	router.HandleFunc("/todo/{id}", routes.DeleteTodoEndPoint).Methods("DELETE")
	router.HandleFunc("/todo/{id}", routes.UpdateTodoEndPoint).Methods("PUT")
	router.HandleFunc("/todo/{id}/timespent", routes.TimeSpentEndPoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))

}
