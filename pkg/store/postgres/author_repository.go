package postgres

import (
	"GO_RESTful_API/pkg/model"
	"context"
)

type AuthorRepository struct {
	store *PostgresStore
	table string
}

func NewAuthorRepository(db *PostgresStore) *AuthorRepository {
	return &AuthorRepository{
		store: db,
		table: AuthorTable,
	}
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, entity model.Author) (bool, error)
func (r *AuthorRepository) DeleteAuthor(ID string) (bool, error)
func (r *AuthorRepository) UpdateAuthor(ID string, entity model.Author) (bool, error)
func (r *AuthorRepository) GetAuthorByID(ID string) (model.Author, error)
func (r *AuthorRepository) GetAllAuthors() ([]model.Author, error)
func (r *AuthorRepository) ExistAuthor(ID string) (bool, error)
func (r *AuthorRepository) ClearAuthors() (bool, error)
