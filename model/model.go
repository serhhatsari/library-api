package model

import "time"

type UserRole string

const (
	Librarian UserRole = "librarian"
	Customer  UserRole = "customer"
)

type Author struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Biography string    `db:"biography"`
	BirthDate time.Time `db:"birth_date"`
}

type Category struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type Book struct {
	ID              int       `db:"id"`
	Title           string    `db:"title"`
	AuthorID        int       `db:"author_id"`
	CategoryID      int       `db:"category_id"`
	PublicationDate time.Time `db:"publication_date"`
	ISBN            string    `db:"isbn"`
	Summary         string    `db:"summary"`
	CoverImageURL   string    `db:"cover_image_url"`
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
