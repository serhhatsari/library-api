package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
	"github.com/serhhatsari/library-api/repository"
	"io"
	"net/http"
	"strconv"
)

func GetBooks(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		books, err := repo.GetBooks()

		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(books) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := jsoniter.Marshal(books)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(json)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func CreateBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &book)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.CreateBook(&book)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strId := chi.URLParam(r, "id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		book, err := repo.GetBookByID(id)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonBook, err := jsoniter.Marshal(book)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(jsonBook)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var book model.Book
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &book)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.UpdateBook(id, &book)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteBook(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Logger().Error(err.Error())
			if err.Error() == "no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.DeleteBook(id)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
