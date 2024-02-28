package controller

import (
	"GO_RESTful_API/api"
	"GO_RESTful_API/pkg/mapper/book"
	"GO_RESTful_API/pkg/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type AuthorController struct {
	service services.Service
}

func NewAuthorController(s services.Service) *AuthorController {
	return &AuthorController{
		service: s,
	}
}

func (ac *AuthorController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := ac.service.GetAll()
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *AuthorController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	res, err := ac.service.GetByID(id)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (ac *AuthorController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	api.ErrorResponse(&w, err)
	author, err := book.MapJsonToAuthor(body)
	api.ErrorResponse(&w, err)
	ok, err := ac.service.Create(&author)
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusCreated)
	}
}

func (ac *AuthorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	api.ErrorResponse(&w, err)
	author, err := book.MapJsonToAuthor(body)
	api.ErrorResponse(&w, err)
	ok, err := ac.service.Update(id, &author)
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusCreated)
	}

}

func (ac *AuthorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ok, err := ac.service.Delete(id)
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (ac *AuthorController) Clear(w http.ResponseWriter, _ *http.Request) {
	ok, err := ac.service.Clear()
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (ac *AuthorController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/author", ac.GetAll).Methods("GET")
	router.HandleFunc("/author/{id}", ac.Get).Methods("GET")
	router.HandleFunc("/author", ac.Create).Methods("POST")
	router.HandleFunc("/author/{id}", ac.Update).Methods("PUT")
	router.HandleFunc("/author/{id}", ac.Delete).Methods("DELETE")
	router.HandleFunc("/author", ac.Clear).Methods("DELETE")
}
