package handler

import (
	"encoding/json"
	"github.com/serhhatsari/library-api/repository"
	"net/http"
)

func GetUsers(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func CreateUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func GetUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func UpdateUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}

func DeleteUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]byte{})
		if err != nil {
			return
		}
	}
}
