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
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database
	config.ConnectDB()

	// Set up routes
	router := routes.SetupRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
