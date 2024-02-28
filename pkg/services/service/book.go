package service

import (
	"GO_RESTful_API/pkg/entities"
	"GO_RESTful_API/pkg/errors"
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

func (bs BookService) Create(entity entities.Entity) (bool, error) {
	logger.Log("trace", "Book was received from http for creation.")
	logger.Log("debug", fmt.Sprintf("Book: %s", entity))
	var ok = false
	var err error = nil
	if bs.validator.Valid(entity) {
		ok, err = bs.repo.Create(entity)
		return ok, err
	}

	jsonString, err := bs.validator.GetJsonErrors()
	if err != nil {
		return false, err
	}
	err = errors.NewError("Validation Error", jsonString, nil)
	logger.Log("trace", "The creation of the Book was finished on service layer.")
	logger.Log("warning", fmt.Sprintf("Validation errors. %s", bs.validator.GetErrors()))

	return false, err
}

func (bs BookService) Delete(ID string) (bool, error) {
	logger.Log("trace", "Book UUID was received from http for deleting book.")
	logger.Log("debug", fmt.Sprintf("Book ID: %s", ID))

	return bs.repo.Delete(ID)

}

func (bs BookService) Update(ID string, entity entities.Entity) (bool, error) {
	logger.Log("trace", "Book UUID and updated book entities was received from http for updating book.")
	logger.Log("debug", fmt.Sprintf("Book ID: %s \t New Book: %s", ID, entity))

	var ok = false
	var err error = nil
	if bs.validator.Valid(entity) {
		ok, err = bs.repo.Update(ID, entity)
		return ok, err
	}
	jsonString, err := bs.validator.GetJsonErrors()
	if err != nil {
		return false, err
	}
	err = errors.NewError("Validation Error", jsonString, nil)
	logger.Log("trace", "The updating of the book was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation errors. %s", bs.validator.GetErrors()))
	return false, err
}

func (bs BookService) GetByID(ID string) (entities.Entity, error) {
	logger.Log("trace", "Book UUID was received from http for receiving book.")
	logger.Log("debug", fmt.Sprintf("Book ID: %s", ID))
	return bs.repo.GetByID(ID)
}

func (bs BookService) GetAll() ([]entities.Entity, error) {
	logger.Log("trace", "Receiving all books.")
	return bs.repo.GetAll()
}

func (bs BookService) Clear() (bool, error) {
	logger.Log("trace", "Book UUID was received from http for receiving book.")
	return bs.repo.Clear()
}
