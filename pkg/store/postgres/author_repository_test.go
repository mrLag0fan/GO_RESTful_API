package postgres_test

import (
	"GO_RESTful_API/pkg/store/postgres"
	"context"
)

func (s *StoreSuite) TestAuthorRepository_Create() {
	author := postgres.TestAuthor(s.T())

	_, err := s.DB.CreateAuthor(context.Background(), author)
	s.NoError(err)

	actual, err := s.DB.GetAuthorByID(author.ID)
	s.NoError(err)

	actual.Birthdate = actual.Birthdate.UTC()

	s.Equal(author, actual)
}
