package repository

import (
	"context"

	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
)

func (repo *Repository) GetBooks() ([]model.Book, error) {
	var books []model.Book

	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM books")
	if err != nil {
		log.Logger().Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.CategoryID, &book.PublicationDate, &book.ISBN, &book.Summary, &book.CoverImageURL)
		if err != nil {
			log.Logger().Error(err.Error())
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (repo *Repository) GetBookByID(id int) (model.Book, error) {
	var book model.Book

	row := repo.conn.QueryRow(context.Background(), "SELECT * FROM books WHERE id = $1", id)

	err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.CategoryID, &book.PublicationDate, &book.ISBN, &book.Summary, &book.CoverImageURL)
	if err != nil {
		log.Logger().Error(err.Error())
		return model.Book{}, err
	}

	return book, nil
}

func (repo *Repository) CreateBook(book *model.Book) error {

	_, err := repo.conn.Exec(context.Background(), "INSERT INTO books (title, author_id, category_id, publication_date, isbn, summary, cover_image_url) VALUES ($1, $2, $3, $4, $5, $6, $7)", book.Title, book.AuthorID, book.CategoryID, book.PublicationDate, book.ISBN, book.Summary, book.CoverImageURL)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) UpdateBook(id int, book *model.Book) error {

	_, err := repo.conn.Exec(context.Background(), "UPDATE books SET title = $1, author_id = $2, category_id = $3, publication_date = $4, isbn = $5, summary = $6, cover_image_url = $7 WHERE id = $8", book.Title, book.AuthorID, book.CategoryID, book.PublicationDate, book.ISBN, book.Summary, book.CoverImageURL, id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) DeleteBook(id int) error {

	_, err := repo.conn.Exec(context.Background(), "DELETE FROM books WHERE id = $1", id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}
