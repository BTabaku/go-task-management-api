package models

import "gorm.io/gorm"

// Task represents a task in the task management system.
type Task struct {
	gorm.Model
	Title       string `json:"title"`       // The title of the task
	Description string `json:"description"` // The description of the task
	Status      string `json:"status"`      // The status of the task
}
