package repository

import (
	"context"

	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
)

func (repo *Repository) GetLoans() ([]model.BookLoan, error) {
	var bookLoans []model.BookLoan

	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM book_loans")
	if err != nil {
		log.Logger().Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		var bookLoan model.BookLoan
		err = rows.Scan(&bookLoan.ID, &bookLoan.BookID, &bookLoan.UserID, &bookLoan.DueDate, &bookLoan.ReturnedDate)
		if err != nil {
			log.Logger().Error(err.Error())
			return nil, err
		}
		bookLoans = append(bookLoans, bookLoan)
	}

	return bookLoans, nil
}

func (repo *Repository) GetLoanByID(id int) (model.BookLoan, error) {
	var bookLoan model.BookLoan

	row := repo.conn.QueryRow(context.Background(), "SELECT * FROM book_loans WHERE id = $1", id)

	err := row.Scan(&bookLoan.ID, &bookLoan.BookID, &bookLoan.UserID, &bookLoan.DueDate, &bookLoan.ReturnedDate)
	if err != nil {
		log.Logger().Error(err.Error())
		return model.BookLoan{}, err
	}

	return bookLoan, nil
}

func (repo *Repository) CreateLoan(bookLoan *model.BookLoan) error {

	_, err := repo.conn.Exec(context.Background(), "INSERT INTO book_loans (book_id, user_id, due_date, returned_date) VALUES ($1, $2, $3, $4)", bookLoan.BookID, bookLoan.UserID, bookLoan.DueDate, bookLoan.ReturnedDate)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) UpdateLoan(id int, bookLoan *model.BookLoan) error {

	_, err := repo.conn.Exec(context.Background(), "UPDATE book_loans SET book_id = $1, user_id = $2, due_date = $3, returned_date = $4 WHERE id = $5", bookLoan.BookID, bookLoan.UserID, bookLoan.DueDate, bookLoan.ReturnedDate, id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}

func (repo *Repository) DeleteLoan(id int) error {

	_, err := repo.conn.Exec(context.Background(), "DELETE FROM book_loans WHERE id = $1", id)

	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	return nil
}
