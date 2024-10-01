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
	_, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	config.InitDatabase()

	// Register routes
	router := routes.RegisterRoutes()

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
