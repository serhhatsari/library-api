package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/serhhatsari/library-api/app"
	log "github.com/serhhatsari/library-api/pkg/logger"
	"github.com/serhhatsari/library-api/repository"
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
		log.Logger().Fatal(err.Error())
		return
	}

}
