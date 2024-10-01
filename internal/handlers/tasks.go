package handlers

import (
	"encoding/json"
	"go-task-management-api/internal/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	tasks []models.Task
	mu    sync.Mutex
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
