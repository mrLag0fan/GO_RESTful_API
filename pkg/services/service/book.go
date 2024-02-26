package service

import (
	"GO_RESTful_API/pkg/entities"
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/repository"
	"GO_RESTful_API/pkg/validation"
	"fmt"
)

type BookService struct {
	repo      repository.Repository
	validator validation.Validator
}

func NewBookService(repo repository.Repository, validator validation.Validator) *BookService {
	logger.Log("trace", "The new book service was created.")
	return &BookService{repo: repo, validator: validator}
}

func (a BookService) Create(entity entities.Entity) (bool, map[string]string) {
	logger.Log("trace", "Book was received from http for creation.")
	logger.Log("debug", fmt.Sprintf("Book: %s", entity))

	if a.validator.Valid(entity) {
		a.repo.Create(entity)
		return true, a.validator.GetErrors()
	}

	logger.Log("trace", "The creation of the Book was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation error. %s", a.validator.GetErrors()))

	return false, a.validator.GetErrors()

}

func (a BookService) Delete(ID string) bool {
	logger.Log("trace", "Book UUID was received from http for deleting book.")
	logger.Log("debug", fmt.Sprintf("Book ID: %s", ID))

	return a.repo.Delete(ID)

}

func (a BookService) Update(ID string, entity entities.Entity) (bool, map[string]string) {
	logger.Log("trace", "Book UUID and updated book entities was received from http for updating book.")
	logger.Log("debug", fmt.Sprintf("Book ID: %s \t New Book: %s", ID, entity))

	if a.validator.Valid(entity) {
		a.repo.Update(ID, entity)
	}

	logger.Log("trace", "The updating of the book was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation error. %s", a.validator.GetErrors()))
	return false, a.validator.GetErrors()
}

func (a BookService) GetByID(ID string) entities.Entity {
	logger.Log("trace", "Book UUID was received from http for receiving book.")
	logger.Log("debug", fmt.Sprintf("Book ID: %s", ID))
	return a.repo.GetByID(ID)
}

func (a BookService) GetAll() []entities.Entity {
	logger.Log("trace", "Receiving all books.")
	return a.repo.GetAll()
}

func (a BookService) Clear() bool {
	logger.Log("trace", "Book UUID was received from http for receiving book.")
	return a.repo.Clear()
}
