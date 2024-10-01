package handlers

import (
	"encoding/json"
	"go-task-management-api/internal/config"
	"go-task-management-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	config.DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the URL parameters from the request
	params := mux.Vars(r)

	// Convert the "id" parameter from a string to an integer
	id, _ := strconv.Atoi(params["id"])

	// Delete the task with the given ID from the database
	config.DB.Delete(&models.Task{}, id)

	// Set the HTTP status code to 204 No Content to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}
