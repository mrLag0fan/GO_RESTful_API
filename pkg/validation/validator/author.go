package validator

import (
	e "GO_RESTful_API/pkg/entities"
	"GO_RESTful_API/pkg/entities/impl/book"
	"GO_RESTful_API/pkg/errors"
	"GO_RESTful_API/pkg/logger"
	"encoding/json"
	"regexp"
)

type AuthorValidator struct {
	err map[string]string
}

func NewAuthorValidator() *AuthorValidator {
	return &AuthorValidator{make(map[string]string)}
}

func (validator *AuthorValidator) Valid(entity e.Entity) bool {
	logger.Log("trace", "Author validation started....")
	validator.err = make(map[string]string)
	author, ok := entity.(*book.Author)
	if !ok {
		validator.err["Wrong entity"] = "Passed entities is not of type Author"
		return false
	}
	validator.validName(*author)
	validator.validSurname(*author)
	validator.validDeathDate(*author)
	logger.Log("trace", "Author validation finished.")
	return !(len(validator.err) > 0)
}

func (validator *AuthorValidator) GetErrors() map[string]string {
	return validator.err
}

func (validator *AuthorValidator) GetJsonErrors() (string, error) {
	errJson, err := json.Marshal(validator.GetErrors())
	if err != nil {
		err = errors.NewError("Json Format Error", err.Error(), &err)
		return err.Error(), err
	}
	return string(errJson), nil
}

func (validator *AuthorValidator) validName(entity book.Author) {
	logger.Log("trace", "Author name validation started....")
	if entity.Name == "" {
		validator.err["Author Name Length"] = "Author name shouldn't be empty string."
	}
	regex := regexp.MustCompile("^[A-Za-z]+$")
	if !regex.MatchString(entity.Name) {
		validator.err["Author Name Only Letters"] = "Author name should only contains letters."
	}
	logger.Log("trace", "Author name validation finished.")
}

func (validator *AuthorValidator) validSurname(entity book.Author) {
	logger.Log("trace", "Author surname validation started....")
	if entity.Name == "" {
		validator.err["Author Surname Length"] = "Author surname shouldn't be empty string."
	}
	regex := regexp.MustCompile("^[A-Za-z]+$")
	if !regex.MatchString(entity.Name) {
		validator.err["Author Surname Only Letters"] = "Author surname should only contains letters."
	}
	logger.Log("trace", "Author surname validation finished.")
}

func (validator *AuthorValidator) validDeathDate(entity book.Author) {
	logger.Log("trace", "Author surname validation started....")
	if entity.DeathDate.Before(entity.Birthdate) {
		validator.err["Author Death Date"] = "Author Death Date should be after birth date."
	}
	logger.Log("trace", "Author death date validation finished.")
}
