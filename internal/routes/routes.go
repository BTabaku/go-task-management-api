package routes

import (
	"go-task-management-api/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")

	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	return router

}
