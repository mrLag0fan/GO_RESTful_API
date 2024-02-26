package validator

import (
	e "GO_RESTful_API/pkg/entities"
	"GO_RESTful_API/pkg/entities/impl/book"
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/repository"
	"fmt"
)

type BookValidator struct {
	err  map[string]string
	repo repository.Repository
}

func NewBookValidator(repository1 repository.Repository) *BookValidator {
	return &BookValidator{err: make(map[string]string), repo: repository1}
}

func (validator *BookValidator) Valid(entity e.Entity) bool {
	logger.Log("trace", "Book validation started....")
	validator.err = make(map[string]string)
	author, ok := entity.(*book.Book)
	if !ok {
		validator.err["Wrong entity"] = "Passed entities is not of type Book"
		return false
	}
	validator.validTitle(*author)
	validator.validNumberOfPages(*author)
	validator.validAuthorID(*author)
	logger.Log("trace", "Book validation finished.")
	return !(len(validator.err) > 0)
}

func (validator *BookValidator) GetErrors() map[string]string {
	return validator.err
}

func (validator *BookValidator) validTitle(entity book.Book) {
	logger.Log("trace", "Book title validation started....")
	if entity.Title == "" {
		validator.err["Book Name Length"] = "Book title shouldn't be empty string."
	}
	logger.Log("trace", "Book title validation finished.")
}

func (validator *BookValidator) validNumberOfPages(entity book.Book) {
	logger.Log("trace", "Book number of pages validation started....")
	if entity.NumberOfPages <= 0 {
		validator.err["Book Death Date"] = "Book number of pages should be >= 0"
	}
	logger.Log("trace", "Book number of pages validation finished.")
}

func (validator *BookValidator) validAuthorID(entity book.Book) {
	logger.Log("trace", "Book number of pages validation started....")
	if !validator.repo.Exist(entity.AuthorID) {
		validator.err["Book Author ID"] = fmt.Sprintf("Book author with %s ID dosen't exist", entity.AuthorID)
	}
	logger.Log("trace", "Book number of pages validation finished.")
}
