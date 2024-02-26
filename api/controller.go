package api

import (
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
