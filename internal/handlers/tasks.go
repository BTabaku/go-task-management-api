package handlers

import (
	"encoding/json"
	"go-task-management-api/internal/models"
	"net/http"
	"strconv" // This package provides functions to convert strings to other data types
	"sync"

	"github.com/gorilla/mux"
)

var (
	tasks []models.Task
	mu    sync.Mutex // this mutex will be used to synchronize access to the tasks slice
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock() // this will ensure that the mutex is unlocked when the function returns

	w.Header().Set("Content-Type", "application/json") // Set the Content-Type header to application/json
	json.NewEncoder(w).Encode(tasks)                   // Encode the tasks slice to JSON and write it to the response writer
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task) // Decode the request body into a Task struct
	task.ID = len(tasks) + 1              // Set the ID of the task
	tasks = append(tasks, task)           // Append the task to the tasks slice
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task) // Encode the task to JSON and write it to the response writer
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get the URL parameters
	id, _ := strconv.Atoi(params["id"])
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
