package controller

import (
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/services"
	"encoding/json"
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

func (bc *BookController) GetAll(w http.ResponseWriter, r *http.Request) {
	res := bc.service.GetAll()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Log("errors", err.Error())
	}
}

func (bc *BookController) Get(w http.ResponseWriter, r *http.Request) {

}

func (bc *BookController) Create(w http.ResponseWriter, r *http.Request) {}

func (bc *BookController) Update(w http.ResponseWriter, r *http.Request) {}

func (bc *BookController) Delete(w http.ResponseWriter, r *http.Request) {}

func (bc *BookController) Clear(w http.ResponseWriter, r *http.Request) {}

func (bc *BookController) HandleRequests() {
	http.HandleFunc("/book", bc.GetAll)
}
