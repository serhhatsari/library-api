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

func GetLoans(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loans, err := repo.GetLoans()
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(loans) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		res, err := jsoniter.Marshal(loans)
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

func CreateLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookLoan model.BookLoan

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = jsoniter.Unmarshal(body, &bookLoan)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = repo.CreateLoan(&bookLoan)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		loan, err := repo.GetLoanByID(id)
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

		res, err := jsoniter.Marshal(loan)
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

func UpdateLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var loan model.BookLoan

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

		err = jsoniter.Unmarshal(body, &loan)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = repo.UpdateLoan(id, &loan)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteLoan(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = repo.DeleteLoan(id)
		if err != nil {
			log.Logger().Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
