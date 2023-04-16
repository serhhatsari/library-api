package repository

import (
	"context"

	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
)

func (repo *Repository) GetCategories() ([]model.Category, error) {
	var categories []model.Category

	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM categories")
	if err != nil {
		log.Logger().Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			log.Logger().Error(err.Error())
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repo *Repository) GetCategoryByID(id int) (model.Category, error) {
	var category model.Category

	row := repo.conn.QueryRow(context.Background(), "SELECT * FROM categories WHERE id = $1", id)

	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		log.Logger().Error(err.Error())
		return model.Category{}, err
	}

	return category, nil
}

func (repo *Repository) CreateCategory(category *model.Category) error {

	_, err := repo.conn.Exec(context.Background(), "INSERT INTO categories (name, description) VALUES ($1, $2)", category.Name, category.Description)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) UpdateCategory(id int, category *model.Category) error {

	_, err := repo.conn.Exec(context.Background(), "UPDATE categories SET name = $1, description = $2 WHERE id = $3", category.Name, category.Description, id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) DeleteCategory(id int) error {

	_, err := repo.conn.Exec(context.Background(), "DELETE FROM categories WHERE id = $1", id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}
