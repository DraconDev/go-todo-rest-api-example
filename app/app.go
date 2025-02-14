package app

import (
	"encoding/json"
	"net/http"
)

// SetupRouter configures and returns a new HTTP handler.
func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	// A simple check endpoint
	mux.HandleFunc("/", helloHandler)
	// Endpoint to return a list of todos
	mux.HandleFunc("/todos", todosHandler)
	return mux
}

// helloHandler sends a welcome message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to the Todo REST API cause why not"))
}

// Todo represents a simple task structure.
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

// todosHandler returns a JSON array of todos.
func todosHandler(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{ID: 1, Task: "Write code", Done: false},
		{ID: 2, Task: "Test code", Done: false},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}