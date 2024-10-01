// This is the entry point of your application. It sets up the server, initializes routes, and starts listening for incoming HTTP requests.

package main

import (
	"go-task-management-api/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
