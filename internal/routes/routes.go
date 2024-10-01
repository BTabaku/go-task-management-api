package routes

import (
	"go-task-management-api/internal/handlers"
	"go-task-management-api/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")

	// Protected routes
	router.Handle("/tasks", middleware.Authenticate(http.HandlerFunc(handlers.CreateTask))).Methods("POST")
	router.Handle("/tasks/{id}", middleware.Authenticate(http.HandlerFunc(handlers.DeleteTask))).Methods("DELETE")

	return router
}
