package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"` 
    Email    string `json:"email"` 
    Password string `json:"password"`
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get all users")
	}
	
	func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get a user")
	}
	
	func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a user")
	}
	
	func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a user")
	}
	
	func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a user")
	}


func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(context.Background())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":8080", nil)
}
