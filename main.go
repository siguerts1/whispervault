// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siguerts1/whispervault/api"
	"github.com/siguerts1/whispervault/database"
)

func main() {
	router := mux.NewRouter()

	// Initialize the database connection
	err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing the database:", err)
	}

	// Use the authenticate middleware for all routes
	router.Use(api.Authenticate(api.GetToken()))

	// Define your routes here
	router.HandleFunc("/secret", api.PushSecret).Methods("POST")
	router.HandleFunc("/secret/{key}", api.GetSecret).Methods("GET")

	port := ":8080" // Default port
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

