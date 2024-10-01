package main

import (
	"log"
	"net/http"
	"os"
	"task-management-api/internal/routes"
)

func main() {
	router := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
