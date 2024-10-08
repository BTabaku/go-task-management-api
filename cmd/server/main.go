package main

import (
	"go-task-management-api/internal/config"
	"go-task-management-api/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
