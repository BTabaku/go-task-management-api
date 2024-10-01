
### [`README.md`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/README.md")

# Task Management API

This is a simple task management API built with Go, Gorilla Mux, and GORM.

## Getting Started

### Prerequisites

- Go 1.22.5 or later
- Docker
- Docker Compose

### Project Structure

```
cmd/
	server/
		main.go
docker-compose.yml
Dockerfile
go.mod
go.sum
internal/
	config/
		config.go
	handlers/
		tasks.go
	middleware/
		auth.go
	models/
		tasks.go
	routes/
		routes.go
README.md
```

### Running the Application

1. **Build and run the application using Docker Compose:**

   ```sh
   docker-compose up --build
   ```

2. The API will be available at [`http://localhost:8080`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Fcmd%2Fserver%2Fmain.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A13%2C%22character%22%3A11%7D%7D%2C%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fhandlers%2Ftasks.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A45%2C%22character%22%3A15%7D%7D%2C%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fmiddleware%2Fauth.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A10%2C%22character%22%3A8%7D%7D%5D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "Go to definition").

### API Endpoints

- `GET /tasks` - Retrieve all tasks
- `POST /tasks` - Create a new task
- `DELETE /tasks/{id}` - Delete a task by ID

### Example Request

**Create a new task:**

```sh
curl -X POST http://localhost:8080/tasks -d '{"title":"New Task","description":"Task description","status":"Pending"}' -H "Content-Type: application/json"
```

### Authentication Middleware

The authentication middleware is defined in [`internal/middleware/auth.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fmiddleware%2Fauth.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/internal/middleware/auth.go"). It uses JWT for authentication.

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

The database is set up using GORM. The [`Task`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fhandlers%2Ftasks.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A13%2C%22character%22%3A16%7D%7D%2C%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fmodels%2Ftasks.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A4%2C%22character%22%3A3%7D%7D%5D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "Go to definition") model is defined in [`internal/models/tasks.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fmodels%2Ftasks.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/internal/models/tasks.go").

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

The handlers for CRUD operations are defined in [`internal/handlers/tasks.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fhandlers%2Ftasks.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/internal/handlers/tasks.go").

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

The routes are registered in [`internal/routes/routes.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Froutes%2Froutes.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/internal/routes/routes.go").

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

The main entry point of the application is in [`cmd/server/main.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Fcmd%2Fserver%2Fmain.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/cmd/server/main.go").

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
   - Apply the [`Authenticate`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Finternal%2Fmiddleware%2Fauth.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A9%2C%22character%22%3A5%7D%7D%5D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "Go to definition") middleware to routes that require authentication.

3. **Write unit tests for your handlers:**
   - Create a new file `internal/handlers/tasks_test.go` and write tests for each handler.

4. **Dockerize the application:**
   - Ensure the [`Dockerfile`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2FDockerfile%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/Dockerfile") and [`docker-compose.yml`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2Fdocker-compose.yml%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/docker-compose.yml") are correctly set up to build and run the application.

5. **Update the documentation:**
   - Add more detailed usage instructions and examples in the [`README.md`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2FBTABAKU%2Fdev%2Fgolang%2Fgo-task-management-api%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%227a4f1405-a3ea-41b4-83ee-63c25e85267d%22%5D "/home/BTABAKU/dev/golang/go-task-management-api/README.md").