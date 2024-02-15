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
	author := e.EntityToAuthor(entity)
	_, err := repo.DB.Exec(`INSERT INTO "author" VALUES ($1, $2, $3, $4, $5)`,
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
	_, err := repo.DB.Exec(`DELETE FROM author WHERE id = $1`, ID)
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	return true
}

func (repo *AuthorRepository) Update(ID string, entity e.Entity) bool {
	author := e.EntityToAuthor(entity)
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
		log.Fatalf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	return false
}

func (repo *AuthorRepository) GetByID(ID string) e.Entity {
	row := repo.DB.QueryRow("SELECT * FROM author WHERE id = $1", ID)
	var author e.Author
	err := row.Scan(&author.ID, &author.Name, &author.Surname, &author.Birthdate, &author.DeathDate)
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return nil
	}
	return &author
}

func (repo *AuthorRepository) GetAll() []e.Entity {
	rows, err := repo.DB.Query("SELECT * FROM author")
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return make([]e.Entity, 0)
	}

	defer rows.Close()

	var res []e.Entity
	for rows.Next() {
		var author e.Author
		err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Surname,
			&author.Birthdate,
			&author.DeathDate)
		if err != nil {
			log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		}
		res = append(res, &author)
	}
	return res
}

func (repo *AuthorRepository) Clear() bool {
	_, err := repo.DB.Exec("DELETE FROM author")
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	return true
}
