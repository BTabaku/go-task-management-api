package handlers

import (
	"bytes"
	"encoding/json"
	"go-task-management-api/internal/config"
	"go-task-management-api/internal/models"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	// Setup
	db = setupTestDB()
	config.SetDB(db)
	// Run tests
	code := m.Run()
	// Teardown
	teardownTestDB(db)
	// Exit
	os.Exit(code)
}

func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=task_manager password=SecureP@ssw0rd! dbname=task_management_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Task{})
	return db
}

func teardownTestDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	sqlDB.Close()
}

func TestCreateTask(t *testing.T) {
	task := models.Task{Title: "Test Task", Description: "Test Description", Status: "Pending"}
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteTask(t *testing.T) {
	// First, create a task to delete
	task := models.Task{Title: "Test Task", Description: "Test Description", Status: "Pending"}
	db.Create(&task)
	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(int(task.ID)), nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
