package handler

import (
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
	"github.com/serhhatsari/library-api/model"
	log "github.com/serhhatsari/library-api/pkg/logger"
	"github.com/serhhatsari/library-api/repository"
	"io"
	"net/http"
	"strconv"
)

func GetAuthors(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authors, err := repo.GetAuthors()
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(authors) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		res, err := jsoniter.Marshal(authors)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(res)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func CreateAuthor(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var author model.Author

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = jsoniter.Unmarshal(body, &author)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.CreateAuthor(author)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAuthor(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		author, err := repo.GetAuthor(id)
		if err != nil {
			log.Logger().Error(err.Error())
			if err.Error() == "no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		res, err := jsoniter.Marshal(author)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(res)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateAuthor(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var author model.Author

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = jsoniter.Unmarshal(body, &author)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = repo.UpdateAuthor(id, author)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteAuthor(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = repo.DeleteAuthor(id)
		if err != nil {
			log.Logger().Error(err.Error())
			if err.Error() == "author not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
