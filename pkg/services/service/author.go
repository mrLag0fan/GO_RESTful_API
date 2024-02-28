package service

import (
	"GO_RESTful_API/pkg/entities"
	"GO_RESTful_API/pkg/errors"
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/repository"
	"GO_RESTful_API/pkg/validation"
	"fmt"
)

type AuthorService struct {
	repo      repository.Repository
	validator validation.Validator
}

func NewAuthorService(repo repository.Repository, validator validation.Validator) *AuthorService {
	logger.Log("trace", "The new author service was created.")
	return &AuthorService{repo: repo, validator: validator}
}

func (as AuthorService) Create(entity entities.Entity) (bool, error) {
	logger.Log("trace", "Author was received from http for creation.")
	logger.Log("debug", fmt.Sprintf("Author: %s", entity))

	var ok = false
	var err error = nil
	if as.validator.Valid(entity) {
		ok, err = as.repo.Create(entity)
		return ok, err
	}

	jsonString, err := as.validator.GetJsonErrors()
	if err != nil {
		return false, err
	}
	err = errors.NewError("Validation Error", jsonString, nil)

	logger.Log("trace", "The creation of the author was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation errors. %s", as.validator.GetErrors()))

	return false, err

}

func (as AuthorService) Delete(ID string) (bool, error) {
	logger.Log("trace", "Author UUID was received from http for deleting author.")
	logger.Log("debug", fmt.Sprintf("Author ID: %s", ID))

	return as.repo.Delete(ID)

}

func (as AuthorService) Update(ID string, entity entities.Entity) (bool, error) {
	logger.Log("trace", "Author UUID and updated author entities was received from http for updating author.")
	logger.Log("debug", fmt.Sprintf("Author ID: %s \t New Author: %s", ID, entity))

	var ok = false
	var err error = nil
	if as.validator.Valid(entity) {
		ok, err = as.repo.Update(ID, entity)
		return ok, err
	}
	jsonString, err := as.validator.GetJsonErrors()
	if err != nil {
		return false, err
	}
	err = errors.NewError("Validation Error", jsonString, nil)

	logger.Log("trace", "The updating of the author was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation errors. %s", as.validator.GetErrors()))
	return false, err
}

func (as AuthorService) GetByID(ID string) (entities.Entity, error) {
	logger.Log("trace", "Author UUID was received from http for receiving author.")
	logger.Log("debug", fmt.Sprintf("Author ID: %s", ID))
	return as.repo.GetByID(ID)
}

func (as AuthorService) GetAll() ([]entities.Entity, error) {
	logger.Log("trace", "Receiving all authors.")
	return as.repo.GetAll()
}

func (as AuthorService) Clear() (bool, error) {
	logger.Log("trace", "Author UUID was received from http for receiving author.")
	return as.repo.Clear()
}
