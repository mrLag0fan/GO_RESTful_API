package store

import (
	"GO_RESTful_API/pkg/model"
	"context"
)

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, entity model.Author) (bool, error)
	DeleteAuthor(ID string) (bool, error)
	UpdateAuthor(ID string, entity model.Author) (bool, error)
	GetAuthorByID(ID string) (model.Author, error)
	GetAllAuthors() ([]model.Author, error)
	ExistAuthor(ID string) (bool, error)
	ClearAuthors() (bool, error)
}
