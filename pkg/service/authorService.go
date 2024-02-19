package service

import (
	"GO_RESTful_API/pkg/entity"
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

func (a AuthorService) Create(entity entity.Entity) (bool, map[string]string) {
	logger.Log("trace", "Author was received from controller for creation.")
	logger.Log("debug", fmt.Sprintf("Author: %s", entity))

	if a.validator.Valid(entity) {
		a.repo.Create(entity)
		return true, a.validator.GetErrors()
	}

	logger.Log("trace", "The creation of the author was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation error. %s", a.validator.GetErrors()))

	return false, a.validator.GetErrors()

}

func (a AuthorService) Delete(ID string) bool {
	logger.Log("trace", "Author UUID was received from controller for deleting author.")
	logger.Log("debug", fmt.Sprintf("Author ID: %s", ID))

	return a.repo.Delete(ID)

}

func (a AuthorService) Update(ID string, entity entity.Entity) (bool, map[string]string) {
	logger.Log("trace", "Author UUID and updated author entity was received from controller for updating author.")
	logger.Log("debug", fmt.Sprintf("Author ID: %s \t New Author: %s", ID, entity))

	if a.validator.Valid(entity) {
		a.repo.Update(ID, entity)
	}

	logger.Log("trace", "The updating of the author was finished on service layer. ")
	logger.Log("warning", fmt.Sprintf("Validation error. %s", a.validator.GetErrors()))
	return false, a.validator.GetErrors()
}

func (a AuthorService) GetByID(ID string) entity.Entity {
	logger.Log("trace", "Author UUID was received from controller for receiving author.")
	logger.Log("debug", fmt.Sprintf("Author ID: %s", ID))
	return a.repo.GetByID(ID)
}

func (a AuthorService) GetAll() []entity.Entity {
	logger.Log("trace", "Receiving all authors.")
	return a.repo.GetAll()
}

func (a AuthorService) Clear() bool {
	logger.Log("trace", "Author UUID was received from controller for receiving author.")
	return a.repo.Clear()
}
