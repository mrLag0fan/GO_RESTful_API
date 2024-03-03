package postgres

import (
	"GO_RESTful_API/pkg/store"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	AuthorTable string = "authors"
)

type PostgresStore struct {
	Conn *pgx.Conn

	store.AuthorRepository
}

// func NewPostgres(cfg config.Postgres) (*PostgresStore, error) {
func NewStore(ctx context.Context) (*PostgresStore, error) {
	// TODO connect ping

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	store := &PostgresStore{
		Conn: conn,
	}

	store.AuthorRepository = NewAuthorRepository(store)

	return store, nil
}
