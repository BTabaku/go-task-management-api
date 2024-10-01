package routes

import (
	"go-task-management-api/internal/handlers"
	"go-task-management-api/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Apply the authentication middleware to the routes
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Apply the authentication middleware to the routes
	router.Use(middleware.Authenticate)

	return router
}
