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
	ID        int      `db:"id"`
	FirstName string   `db:"first_name"`
	LastName  string   `db:"last_name"`
	Email     string   `db:"email"`
	Password  string   `db:"password"`
	Role      UserRole `db:"role"`
}

type BookLoan struct {
	ID           int       `db:"id"`
	BookID       int       `db:"book_id"`
	UserID       int       `db:"user_id"`
	DueDate      time.Time `db:"due_date"`
	ReturnedDate time.Time `db:"returned_date"`
}
