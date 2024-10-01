package main

import (
	"go-task-management-api/internal/config"
	"go-task-management-api/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load configuration
	configFilePath := ".env" // or any other path to your .env file
	_, err := config.LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database
	config.InitDatabase()

	// Register routes
	router := routes.RegisterRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
