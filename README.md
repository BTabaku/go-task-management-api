# Task Management API

This is a simple task management API built with Go, Gorilla Mux, and GORM.

## Getting Started

### Prerequisites

- Go 1.22.5 or later
- Docker
- Docker Compose

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/go-task-management-api.git
   cd go-task-management-api
   ```

2. **Copy the example environment file and update the configuration:**

   ```sh
   cp .env.example .env
   ```

3. **Build and run the application using Docker Compose:**

   ```sh
   docker-compose up --build
   ```

4. The API will be available at [http://localhost:8080](http://localhost:8080).

### Project Structure

```
.
├── cmd
│   └── server
│       └── main.go
├── deployments
│   └── k8s
│       ├── task-api-deployment.yml
│       └── task-api-service.yml
├── docker-compose.yml
├── Dockerfile
├── generate_token.go
├── go.mod
├── go.sum
├── init-db.sh
├── internal
│   ├── config
│   │   ├── config.go
│   │   ├── db.go
│   │   └── db_test.go
│   ├── handlers
│   │   ├── tasks.go
│   │   └── tasks_test.go
│   ├── middleware
│   │   └── auth.go
│   ├── models
│   │   └── tasks.go
│   └── routes
│       └── routes.go
└── README.md

```

### Running the Application

1. **Build and run the application using Docker Compose:**

   ```sh
   docker-compose up --build
   ```

2. The API will be available at [http://localhost:8080](http://localhost:8080).

### API Endpoints

- `GET /tasks` - Retrieve all tasks
- `POST /tasks` - Create a new task
- `DELETE /tasks/{id}` - Delete a task by ID

### Example Request

**Create a new task:**

```sh
curl -X POST http://localhost:8080/tasks \
     -d '{"title":"New Task","description":"Task description","status":"Pending"}' \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### Authentication Middleware





The authentication middleware is defined in `internal/middleware/auth.go`. It uses JWT for authentication.

```go
package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-256-bit-secret"), nil
		})

		if !token.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
```

### Database Setup

The database is set up using GORM. The `Task` model is defined in `internal/models/tasks.go`.

```go
package models

import "gorm.io/gorm"

// Task represents a task in the task management system.
type Task struct {
	gorm.Model
	Title       string `json:"title"`       // The title of the task
	Description string `json:"description"` // The description of the task
	Status      string `json:"status"`      // The status of the task
}
```

### Handlers

The handlers for CRUD operations are defined in `internal/handlers/tasks.go`.

```go
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
```

### Routes

The routes are registered in `internal/routes/routes.go`.

```go
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
```

### Main Entry Point

The main entry point of the application is in `cmd/server/main.go`.

```go
package main

import (
	"go-task-management-api/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
```

### TODO

1. **Implement the database logic in your handlers:**
   - Replace the in-memory storage with a database using GORM.

2. **Add authentication middleware to protect your endpoints:**
   - Apply the `Authenticate` middleware to routes that require authentication.

3. **Write unit tests for your handlers:**
   - Create a new file `internal/handlers/tasks_test.go` and write tests for each handler.

4. **Dockerize the application:**
   - Ensure the `Dockerfile` and `docker-compose.yml` are correctly set up to build and run the application.

5. **Update the documentation:**
   - Add more detailed usage instructions and examples in the `README.md`.

### Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
