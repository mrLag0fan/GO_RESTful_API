package repository

import (
	"GO_RESTful_API/pkg/database"
	e "GO_RESTful_API/pkg/entity"
	"GO_RESTful_API/pkg/logger"
	"GO_RESTful_API/pkg/my_uuid"
	"database/sql"
	"log"
	"time"
)

type BookRepository struct {
	DB            *sql.DB
	uuidGenerator my_uuid.UuidGenerator
}

func NewBookRepository(generator my_uuid.UuidGenerator) *BookRepository {
	logger.Log("trace", "The new book repository was created.")
	return &BookRepository{
		DB:            database.DB,
		uuidGenerator: generator,
	}
}

func (repo *BookRepository) Create(entity e.Entity) bool {
	logger.Log("trace", "Book insertion into database begun....")
	book := e.EntityToBook(entity)
	_, err := repo.DB.Exec(`INSERT INTO "book" VALUES ($1, $2, $3, $4, $5)`,
		repo.uuidGenerator.GenerateUUID(),
		book.Title,
		book.NumberOfPages,
		book.Description,
		book.AuthorID)
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Book insertion into database finished.")
	return true
}

func (repo *BookRepository) Delete(ID string) bool {
	logger.Log("trace", "Book deleting from database begun....")
	_, err := repo.DB.Exec(`DELETE FROM book WHERE id = $1`, ID)
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	logger.Log("trace", "Book deleting from database finished.")
	return true
}

func (repo *BookRepository) Update(ID string, entity e.Entity) bool {
	logger.Log("trace", "Book updating from database begun....")
	book := e.EntityToBook(entity)
	_, err := repo.DB.Exec(`UPDATE "book" SET 
                    title=$1, 
                    numberofpages=$2,
                    description=$3, 
                    authorid=$4
                    WHERE id=$6`,
		book.Title,
		book.NumberOfPages,
		book.Description,
		book.AuthorID,
		ID)
	if err != nil {
		logger.Log("error", err.Error())
		return false
	}
	logger.Log("trace", "Book updating from database finished.")
	return false
}

func (repo *BookRepository) GetByID(ID string) e.Entity {
	logger.Log("trace", "Book receiving from database begun....")
	row := repo.DB.QueryRow("SELECT * FROM book WHERE id = $1", ID)
	var book e.Book
	err := row.Scan(&book.ID, &book.Title, &book.NumberOfPages, &book.Description, &book.AuthorID)
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return nil
	}
	logger.Log("trace", "Book receiving from database finished.")
	return &book
}

func (repo *BookRepository) GetAll() []e.Entity {
	logger.Log("trace", "Receiving all books from database begun....")
	rows, err := repo.DB.Query("SELECT * FROM book")
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return make([]e.Entity, 0)
	}

	defer rows.Close()

	var res []e.Entity
	for rows.Next() {
		var book e.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.NumberOfPages,
			&book.Description,
			&book.AuthorID)
		if err != nil {
			log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		}
		res = append(res, &book)
	}
	logger.Log("trace", "Receiving all books from database finished.")
	return res
}

func (repo *BookRepository) Clear() bool {
	logger.Log("trace", "Clearing books from database begun....")
	_, err := repo.DB.Exec("DELETE FROM book")
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false
	}
	logger.Log("trace", "Clearing books from database finished.")
	return true
}
