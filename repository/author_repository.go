package repository

import (
	"context"
	"fmt"
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
			fmt.Println(err)
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}
