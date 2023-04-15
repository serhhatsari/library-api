package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/serhhatsari/library-api/handler"
	"github.com/serhhatsari/library-api/repository"
)

type App struct {
	Repo   *repository.Repository
	Router *chi.Mux
}

func (app *App) Register() {
	// Register Author routes
	app.Router.Get("/authors", handler.GetAuthors(app.Repo))
	app.Router.Post("/authors", handler.CreateAuthor(app.Repo))
	app.Router.Get("/authors/{id}", handler.GetAuthor(app.Repo))
	app.Router.Put("/authors/{id}", handler.UpdateAuthor(app.Repo))
	app.Router.Delete("/authors/{id}", handler.DeleteAuthor(app.Repo))

	// Register Book routes
	app.Router.Get("/books", handler.GetBooks(app.Repo))
	app.Router.Post("/books", handler.CreateBook(app.Repo))
	app.Router.Get("/books/{id}", handler.GetBook(app.Repo))
	app.Router.Put("/books/{id}", handler.UpdateBook(app.Repo))
	app.Router.Delete("/books/{id}", handler.DeleteBook(app.Repo))

	// Register User routes
	app.Router.Get("/users", handler.GetUsers(app.Repo))
	app.Router.Post("/users", handler.CreateUser(app.Repo))
	app.Router.Get("/users/{id}", handler.GetUser(app.Repo))
	app.Router.Put("/users/{id}", handler.UpdateUser(app.Repo))
	app.Router.Delete("/users/{id}", handler.DeleteUser(app.Repo))

	// Register Loan routes
	app.Router.Get("/loans", handler.GetLoans(app.Repo))
	app.Router.Post("/loans", handler.CreateLoan(app.Repo))
	app.Router.Get("/loans/{id}", handler.GetLoan(app.Repo))
	app.Router.Put("/loans/{id}", handler.UpdateLoan(app.Repo))
	app.Router.Delete("/loans/{id}", handler.DeleteLoan(app.Repo))

	// Register Category routes
	app.Router.Get("/categories", handler.GetCategories(app.Repo))
	app.Router.Post("/categories", handler.CreateCategory(app.Repo))
	app.Router.Get("/categories/{id}", handler.GetCategory(app.Repo))
	app.Router.Put("/categories/{id}", handler.UpdateCategory(app.Repo))
	app.Router.Delete("/categories/{id}", handler.DeleteCategory(app.Repo))
}
