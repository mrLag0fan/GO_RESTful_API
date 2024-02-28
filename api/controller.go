package api

import (
	"GO_RESTful_API/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Clear(w http.ResponseWriter, r *http.Request)
	HandleRequests(router *mux.Router)
}

func ErrorResponse(w *http.ResponseWriter, err error) {
	if err != nil {
		logger.Log("error", err.Error())
		http.Error(*w, err.Error(), http.StatusInternalServerError)
	}
}
