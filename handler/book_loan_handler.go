package handler

import (
	"encoding/json"
	"github.com/serhhatsari/library-api/repository"
	"net/http"
)

func GetLoans(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func CreateLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func GetLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func UpdateLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func DeleteLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}
