package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/serhhatsari/library-api/app"
	"github.com/serhhatsari/library-api/repository"
	"net/http"
)

func main() {

	// Connect to the database
	repo := repository.New()
	defer repo.Close()

	// Create the router
	r := chi.NewRouter()

	// Initialize the application
	application := &app.App{
		Repo:   repo,
		Router: r,
	}

	// Register the routes
	application.Register()

	// Start the server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
