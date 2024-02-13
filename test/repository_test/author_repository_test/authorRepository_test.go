package author_repository_test_test

import (
	"GO_RESTful_API/pkg/entity"
	"GO_RESTful_API/pkg/my_uuid/test_uuid"
	"GO_RESTful_API/pkg/repository"
	"testing"
	"time"
)

var repositoryAuthor repository.Repository
var correctAuthor entity.Author
var incorrectAuthor entity.Author

func init() {
	repositoryAuthor = repository.NewAuthorRepository(&test_uuid.TestUUIDGenerator{})
	correctAuthor = entity.Author{
		Name:      "Name",
		Surname:   "Surname",
		Birthdate: time.Now().AddDate(-50, 0, 0),
		DeathDate: time.Now(),
	}
	incorrectAuthor = entity.Author{}
	repositoryAuthor.Clear()
}

func TestCreateCorrect(t *testing.T) {
	got := repositoryAuthor.Create(&correctAuthor)
	want := true
	if got != want {
		t.Errorf("Create(%s) = %t; want %t", correctAuthor, got, want)
	}
}
func TestCreateIncorrect(t *testing.T) {
	got := repositoryAuthor.Create(&incorrectAuthor)
	want := false
	if got != want {
		t.Errorf("Create(%s) = %t; want %t", incorrectAuthor, got, want)
	}
}
