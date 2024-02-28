package main

import (
	"GO_RESTful_API/api"
	"GO_RESTful_API/api/controller"
	"GO_RESTful_API/pkg/my_uuid/impl"
	"GO_RESTful_API/pkg/repository/postgres"
	"GO_RESTful_API/pkg/services/service"
	"GO_RESTful_API/pkg/validation/validator"
)

func main() {
	authorRepo := postgres.NewAuthorRepository(&impl.RealUUIDGenerator{})
	bookRepo := postgres.NewBookRepository(&impl.RealUUIDGenerator{})
	authorValidator := validator.NewAuthorValidator()
	bookValidator := validator.NewBookValidator(authorRepo)
	bookServ := service.NewBookService(bookRepo, bookValidator)
	authorServ := service.NewAuthorService(authorRepo, authorValidator)
	bookController := controller.NewBookController(bookServ)
	authorController := controller.NewAuthorController(authorServ)

	server := api.NewServer([]api.Controller{authorController, bookController})
	server.HandleRequests()
}
