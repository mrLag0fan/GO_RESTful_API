package author

import (
	e "GO_RESTful_API/pkg/entity"
	"regexp"
)

type AuthorValidator struct {
	err map[string]string
}

func NewValidator() *AuthorValidator {
	return &AuthorValidator{make(map[string]string)}
}

func (validator *AuthorValidator) Valid(entity e.Entity) bool {
	author, ok := entity.(*e.Author)
	if !ok {
		validator.err["Wrong entity"] = "Passed entity is not of type Author"
		return false
	}
	validator.validName(*author)
	validator.validSurname(*author)
	validator.validBirthDate(*author)
	validator.validDeathDate(*author)
	return !(len(validator.err) > 0)
}

func (validator *AuthorValidator) validName(entity e.Author) {
	if entity.Name == "" {
		validator.err["Author Name Length"] = "Author name shouldn't be empty string."
	}
	regex := regexp.MustCompile("^[A-Za-z]+$")
	if regex.MatchString(entity.Name) {
		validator.err["Author Name Only Letters"] = "Author name should only contains letters."
	}
}

func (validator *AuthorValidator) validSurname(entity e.Author) {
	if entity.Name == "" {
		validator.err["Author Surname Length"] = "Author surname shouldn't be empty string."
	}
	regex := regexp.MustCompile("^[A-Za-z]+$")
	if regex.MatchString(entity.Name) {
		validator.err["Author Surname Only Letters"] = "Author surname should only contains letters."
	}
}

func (validator *AuthorValidator) validBirthDate(entity e.Author) bool {
	return false
}

func (validator *AuthorValidator) validDeathDate(entity e.Author) bool {
	return false
}
