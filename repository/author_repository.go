package repository

import (
	"context"
	"errors"
	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
)

func (repo *Repository) GetAuthors() ([]model.Author, error) {
	var authors []model.Author

	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM authors")
	if err != nil {
		log.Logger().Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		var author model.Author
		err = rows.Scan(&author.ID, &author.Name, &author.Biography, &author.BirthDate)
		if err != nil {
			log.Logger().Error(err.Error())
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (repo *Repository) CreateAuthor(author model.Author) error {
	_, err := repo.conn.Exec(context.Background(), "INSERT INTO authors (name, biography, birth_date) VALUES ($1, $2, $3)", author.Name, author.Biography, author.BirthDate)
	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) GetAuthor(id int) (model.Author, error) {
	var author model.Author

	row := repo.conn.QueryRow(context.Background(), "SELECT * FROM authors WHERE id = $1", id)
	err := row.Scan(&author.ID, &author.Name, &author.Biography, &author.BirthDate)
	if err != nil {
		log.Logger().Error(err.Error())
		return model.Author{}, err
	}

	return author, nil
}

func (repo *Repository) UpdateAuthor(id int, author model.Author) error {
	_, err := repo.conn.Exec(context.Background(), "UPDATE authors SET name = $1, biography = $2, birth_date = $3 WHERE id = $4", author.Name, author.Biography, author.BirthDate, id)
	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) DeleteAuthor(id int) error {
	tag, err := repo.conn.Exec(context.Background(), "DELETE FROM authors WHERE id = $1", id)
	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("author not found")
	}

	return nil
}
