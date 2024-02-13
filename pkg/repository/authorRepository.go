package repository

import (
	"GO_RESTful_API/pkg/database"
	e "GO_RESTful_API/pkg/entity"
	"GO_RESTful_API/pkg/my_uuid"
	"database/sql"
	"log"
	"time"
)

type AuthorRepository struct {
	DB            *sql.DB
	uuidGenerator my_uuid.UuidGenerator
}

func NewAuthorRepository(generator my_uuid.UuidGenerator) *AuthorRepository {
	return &AuthorRepository{
		DB:            database.DB,
		uuidGenerator: generator,
	}
}

func (repo *AuthorRepository) Create(entity e.Entity) bool {
	author, ok := entity.(*e.Author)
	if !ok {
		log.Fatalf("%s - [ERROR] Passed entity is not of type Author", time.Now())
		return false
	}
	_, err := database.DB.Exec(`INSERT INTO "author" VALUES ($1, $2, $3, $4, $5)`,
		repo.uuidGenerator.GenerateUUID(),
		author.Name,
		author.Surname,
		author.Birthdate,
		author.DeathDate)
	if err != nil {
		log.Fatalf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	return true
}

func (repo *AuthorRepository) Delete(ID string) bool {
	_, err := database.DB.Exec(`DELETE FROM author WHERE id = $1`, ID)
	if err != nil {
		log.Fatalf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	return true
}
func (repo *AuthorRepository) Update(ID string, entity e.Entity) bool {
	return false
}

func (repo *AuthorRepository) GetByID(ID string) *e.Entity {
	return nil
}

func (repo *AuthorRepository) GetAll() []e.Entity {
	return nil
}

func (repo *AuthorRepository) Clear() bool {
	_, err := database.DB.Exec("DELETE FROM author")
	if err != nil {
		log.Fatalf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	return true
}
