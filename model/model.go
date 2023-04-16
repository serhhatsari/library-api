package model

import "time"

type UserRole string

const (
	Librarian UserRole = "librarian"
	Customer  UserRole = "customer"
)

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	AuthorID        int       `json:"author_id"`
	CategoryID      int       `json:"category_id"`
	PublicationDate time.Time `json:"publication_date"`
	ISBN            string    `json:"isbn"`
	Summary         string    `json:"summary"`
	CoverImageURL   string    `json:"cover_image_url"`
}

type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
}

type BookLoan struct {
	ID           int       `json:"id"`
	BookID       int       `json:"book_id"`
	UserID       int       `json:"user_id"`
	DueDate      time.Time `json:"due_date"`
	ReturnedDate time.Time `json:"returned_date"`
}
