package main

import (
	"go-task-management-api/internal/config"
	"go-task-management-api/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize the database
	config.InitDatabase()

	// Register routes
	router := routes.RegisterRoutes()

	// Get the port from the environment variable, default to 9090 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// Start the server
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
