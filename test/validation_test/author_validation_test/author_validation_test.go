package author_validation_test

import (
	"GO_RESTful_API/pkg/entity"
	"GO_RESTful_API/pkg/validation"
	"GO_RESTful_API/pkg/validation/author"
	"testing"
	"time"
)

var correctAuthor entity.Entity
var incorrectAuthor entity.Entity
var validator validation.Validator

func init() {
	correctAuthor = &entity.Author{
		Name:      "Taras",
		Surname:   "Shevchenko",
		Birthdate: time.Date(1814, 3, 9, 0, 0, 0, 0, time.Local),
		DeathDate: time.Date(1861, 3, 10, 0, 0, 0, 0, time.Local),
	}
	incorrectAuthor = &entity.Author{
		Name:      "",
		Surname:   "4564",
		Birthdate: time.Date(1814, 3, 9, 0, 0, 0, 0, time.Local),
		DeathDate: time.Date(900, 3, 10, 0, 0, 0, 0, time.Local),
	}
	validator = &author.AuthorValidator{}
}

func TestAuthorValidatorCorrect(t *testing.T) {
	got := validator.Valid(correctAuthor)
	want := true
	if got != want {
		t.Errorf("Create(%s) = %t; want %t", correctAuthor, got, want)
	}
}

func TestAuthorValidatorIncorrect(t *testing.T) {
	got := validator.Valid(incorrectAuthor)
	want := false
	if got != want {
		t.Errorf("Create(%s) = %t; want %t", correctAuthor, got, want)
	}
}
