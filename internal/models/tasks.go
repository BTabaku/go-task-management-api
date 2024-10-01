package models

// Task represents a task in the task management system.
type Task struct {
	ID          int    `json:"id"`          // The ID of the task
	Title       string `json:"title"`       // The title of the task
	Description string `json:"description"` // The description of the task
	Status      string `json:"status"`      // The status of the task
	CreatedAt   string `json:"created_at"`  // The creation date of the task
	UpdateAt    string `json:"updated_at"`  // The last update date of the task
}
