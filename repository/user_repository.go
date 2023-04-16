package repository

import (
	"context"

	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
)

func (repo *Repository) GetUsers() ([]model.User, error) {
	var users []model.User

	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Logger().Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			log.Logger().Error(err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *Repository) GetUserByID(id int) (model.User, error) {
	var user model.User

	row := repo.conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Logger().Error(err.Error())
		return model.User{}, err
	}

	return user, nil
}

func (repo *Repository) CreateUser(user *model.User) error {

	_, err := repo.conn.Exec(context.Background(), "INSERT INTO users (first_name, last_name, email, password, role) VALUES ($1, $2, $3, $4, $5)", user.FirstName, user.LastName, user.Email, user.Password, user.Role)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) UpdateUser(id int, user *model.User) error {

	_, err := repo.conn.Exec(context.Background(), "UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4, role = $5 WHERE id = $6", user.FirstName, user.LastName, user.Email, user.Password, user.Role, id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) DeleteUser(id int) error {

	_, err := repo.conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}
