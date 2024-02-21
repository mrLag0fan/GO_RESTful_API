package postgres

import (
	"GO_RESTful_API/pkg/database"
	e "GO_RESTful_API/pkg/entities"
	"GO_RESTful_API/pkg/entities/impl/book"
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/my_uuid"
	"database/sql"
)

type AuthorRepository struct {
	DB            *sql.DB
	uuidGenerator my_uuid.UuidGenerator
}

func NewAuthorRepository(generator my_uuid.UuidGenerator) *AuthorRepository {
	logger.Log("trace", "The new author repository was created.")
	return &AuthorRepository{
		DB:            database.DB,
		uuidGenerator: generator,
	}
}

func (repo *AuthorRepository) Create(entity e.Entity) bool {
	logger.Log("trace", "Author insertion into database begun....")
	author := book.EntityToAuthor(entity)
	_, err := repo.DB.Exec(`INSERT INTO "author" VALUES ($1, $2, $3, $4, $5)`,
		repo.uuidGenerator.GenerateUUID(),
		author.Name,
		author.Surname,
		author.Birthdate,
		author.DeathDate)
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Author insertin into database finished.")
	return true
}

func (repo *AuthorRepository) Delete(ID string) bool {
	logger.Log("trace", "Author deleteing from database begun....")
	_, err := repo.DB.Exec(`DELETE FROM author WHERE id = $1`, ID)
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Author deleteing from database finished.")
	return true
}

func (repo *AuthorRepository) Update(ID string, entity e.Entity) bool {
	logger.Log("trace", "Author updating from database begun....")
	author := book.EntityToAuthor(entity)
	_, err := repo.DB.Exec(`UPDATE "author" SET 
                    name=$1, 
                    surname=$2, 
                    birthdate=$3,
                    death_date=$4
                    WHERE id=$6`,
		author.Name,
		author.Surname,
		author.Birthdate,
		author.DeathDate,
		ID)
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Author updating from database finished.")
	return false
}

func (repo *AuthorRepository) GetByID(ID string) e.Entity {
	logger.Log("trace", "Author receiving from database begun....")
	row := repo.DB.QueryRow("SELECT * FROM author WHERE id = $1", ID)
	var author book.Author
	err := row.Scan(&author.ID, &author.Name, &author.Surname, &author.Birthdate, &author.DeathDate)
	if err != nil {
		logger.Log("error", err.Error())
		return nil
	}
	logger.Log("trace", "Author receiving from database finished.")
	return &author
}

func (repo *AuthorRepository) GetAll() []e.Entity {
	logger.Log("trace", "Receiving all authors from database begun....")
	rows, err := repo.DB.Query("SELECT * FROM author")
	if err != nil {
		logger.Log("error", err.Error())
		return make([]e.Entity, 0)
	}

	defer rows.Close()

	var res []e.Entity
	for rows.Next() {
		var author book.Author
		err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Surname,
			&author.Birthdate,
			&author.DeathDate)
		if err != nil {
			logger.Log("error", err.Error())
		}
		res = append(res, &author)
	}
	logger.Log("trace", "Receiving all authors from database finished.")
	return res
}

func (repo *AuthorRepository) Exist(ID string) bool {
	logger.Log("trace", "Checking whether the author exists....")
	var exists bool
	err := repo.DB.QueryRow("SELECT (exists(SELECT 1 FROM author WHERE id = $1))", ID).Scan(&exists)
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Checking whether the author exists finished.")
	return exists
}

func (repo *AuthorRepository) Clear() bool {
	logger.Log("trace", "Clearing authors from database begun....")
	_, err := repo.DB.Exec("DELETE FROM author")
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Clearing authors from database finished.")
	return true
}
