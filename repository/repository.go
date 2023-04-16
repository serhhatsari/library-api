package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func New() *Repository {
	conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	return &Repository{
		conn: conn,
	}
}

func (repo *Repository) Close() {
	repo.conn.Close()
}
