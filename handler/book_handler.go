package handler

import (
	"encoding/json"
	"github.com/serhhatsari/library-api/repository"
	"net/http"
)

func GetBooks(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func CreateBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func GetBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func UpdateBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func DeleteBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}
