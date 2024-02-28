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

type BookController struct {
	service services.Service
}

func NewBookController(s services.Service) *BookController {
	return &BookController{
		service: s,
	}
}

func (bc *BookController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := bc.service.GetAll()
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (bc *BookController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	res, err := bc.service.GetByID(id)
	api.ErrorResponse(&w, err)
	err = json.NewEncoder(w).Encode(res)
	api.ErrorResponse(&w, err)
	w.WriteHeader(http.StatusOK)
}

func (bc *BookController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	api.ErrorResponse(&w, err)
	author, err := book.MapJsonToAuthor(body)
	api.ErrorResponse(&w, err)
	ok, err := bc.service.Create(&author)
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusCreated)
	}
}

func (bc *BookController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	api.ErrorResponse(&w, err)
	author, err := book.MapJsonToAuthor(body)
	api.ErrorResponse(&w, err)
	ok, err := bc.service.Update(id, &author)
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusCreated)
	}
}

func (bc *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ok, err := bc.service.Delete(id)
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (bc *BookController) Clear(w http.ResponseWriter, _ *http.Request) {
	ok, err := bc.service.Clear()
	api.ErrorResponse(&w, err)
	if ok {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (bc *BookController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/book", bc.GetAll).Methods("GET")
	router.HandleFunc("/book/{id}", bc.Get).Methods("GET")
	router.HandleFunc("/book", bc.Create).Methods("POST")
	router.HandleFunc("/book/{id}", bc.Update).Methods("PUT")
	router.HandleFunc("/book/{id}", bc.Delete).Methods("DELETE")
	router.HandleFunc("/book", bc.Clear).Methods("DELETE")
}
