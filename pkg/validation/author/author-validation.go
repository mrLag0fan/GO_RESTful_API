package author

import (
	e "GO_RESTful_API/pkg/entity"
	"GO_RESTful_API/pkg/logger"
	"regexp"
)

type AuthorValidator struct {
	err map[string]string
}

func NewValidator() *AuthorValidator {
	return &AuthorValidator{make(map[string]string)}
}

func (validator *AuthorValidator) Valid(entity e.Entity) bool {
	logger.Log("trace", "Author validation started....")
	validator.err = make(map[string]string)
	author, ok := entity.(*e.Author)
	if !ok {
		validator.err["Wrong entity"] = "Passed entity is not of type Author"
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

func (validator *AuthorValidator) validName(entity e.Author) {
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

func (validator *AuthorValidator) validSurname(entity e.Author) {
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

func (validator *AuthorValidator) validDeathDate(entity e.Author) {
	logger.Log("trace", "Author surname validation started....")
	if entity.DeathDate.Before(entity.Birthdate) {
		validator.err["Author Death Date"] = "Author Death Date should be after birth date."
	}
	logger.Log("trace", "Author death date validation finished.")
}
