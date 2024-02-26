package controller

import (
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/services"
	"encoding/json"
	"github.com/gorilla/mux"
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
	res := ac.service.GetAll()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Log("error", err.Error())
		http.Error(w, "Failed to encode result", http.StatusInternalServerError)
	}
}

func (ac *AuthorController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	res := ac.service.GetByID(id)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Log("error", err.Error())
		http.Error(w, "Failed to encode result", http.StatusInternalServerError)
	}
}

func (ac *AuthorController) Create(w http.ResponseWriter, r *http.Request) {

}

func (ac *AuthorController) Update(w http.ResponseWriter, r *http.Request) {}

func (ac *AuthorController) Delete(w http.ResponseWriter, r *http.Request) {}

func (ac *AuthorController) Clear(w http.ResponseWriter, r *http.Request) {}

func (ac *AuthorController) HandleRequests(router *mux.Router) {
	router.HandleFunc("/author", ac.GetAll).Methods("GET")
	router.HandleFunc("/author/{id}", ac.Get).Methods("GET")
}
