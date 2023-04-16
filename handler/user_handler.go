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

func GetUsers(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := repo.GetUsers()
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		res, err := jsoniter.Marshal(users)
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

func CreateUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user model.User

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = jsoniter.Unmarshal(body, &user)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.CreateUser(&user)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := repo.GetUserByID(id)
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

		res, err := jsoniter.Marshal(user)
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

func UpdateUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user model.User

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

		err = jsoniter.Unmarshal(body, &user)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = repo.UpdateUser(id, &user)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUser(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = repo.DeleteUser(id)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
