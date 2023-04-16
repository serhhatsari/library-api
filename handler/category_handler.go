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

func GetCategories(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		categories, err := repo.GetCategories()

		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(categories) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := jsoniter.Marshal(categories)
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

func CreateCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category model.Category

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &category)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.CreateCategory(&category)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}

func GetCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		strId := chi.URLParam(r, "id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		category, err := repo.GetCategoryByID(id)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonCategory, err := jsoniter.Marshal(category)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(jsonCategory)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}

func UpdateCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var category model.Category
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &category)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.UpdateCategory(id, &category)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteCategory(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.DeleteCategory(id)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
