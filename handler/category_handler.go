package handler

import (
	"encoding/json"
	"github.com/serhhatsari/library-api/repository"
	"net/http"
)

func GetCategories(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func CreateCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func GetCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func UpdateCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func DeleteCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}
