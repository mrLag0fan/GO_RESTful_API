package controller

import (
	"GO_RESTful_API/pkg/logger"
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

func (ac *AuthorController) GetAll(w http.ResponseWriter, r *http.Request) {
	res, _ := ac.service.GetAll()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Log("errors", err.Error())
		http.Error(w, "Failed to encode result", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (ac *AuthorController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	res, _ := ac.service.GetByID(id)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Log("errors", err.Error())
		http.Error(w, "Failed to encode result", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (ac *AuthorController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	author, err := book.MapJsonToAuthor(body)
	if err != nil {
		logger.Log("errors", err.Error())
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	ac.service.Create(&author)
	w.WriteHeader(http.StatusCreated)
}

func (ac *AuthorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	author, err := book.MapJsonToAuthor(body)
	if err != nil {
		logger.Log("errors", err.Error())
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	ac.service.Update(id, &author)
	w.WriteHeader(http.StatusCreated)
}

func (ac *AuthorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ok, _ := ac.service.Delete(id)
	if ok {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (ac *AuthorController) Clear(w http.ResponseWriter, r *http.Request) {
	ac.service.Clear()
}

func (ac *AuthorController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/author", ac.GetAll).Methods("GET")
	router.HandleFunc("/author/{id}", ac.Get).Methods("GET")
	router.HandleFunc("/author", ac.Create).Methods("POST")
	router.HandleFunc("/author/{id}", ac.Update).Methods("PUT")
	router.HandleFunc("/author/{id}", ac.Delete).Methods("DELETE")
	router.HandleFunc("/author", ac.Clear).Methods("DELETE")
}
