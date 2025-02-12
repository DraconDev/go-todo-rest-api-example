package main

import (
	"log"
	"net/http"

	"github.com/mingrammer/go-todo-rest-api-example/app"
	"github.com/mingrammer/go-todo-rest-api-example/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize the application (router, handlers, etc)
	router := app.SetupRouter()

	log.Printf("Server is running on port %s", cfg.Port)
	// Start the HTTP server
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}