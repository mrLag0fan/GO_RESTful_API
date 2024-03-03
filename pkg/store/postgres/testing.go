package postgres

import (
	"GO_RESTful_API/pkg/model"
	"testing"
	"time"
)

func TestAuthor(t *testing.T) model.Author {
	t.Helper()

	return model.Author{
		ID:        "123",
		Name:      "123",
		Surname:   "123",
		Birthdate: time.Now(),
		DeathDate: time.Now(),
	}
}
