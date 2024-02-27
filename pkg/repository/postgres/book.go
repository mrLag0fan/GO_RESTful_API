package postgres

import (
	"GO_RESTful_API/pkg/database"
	e "GO_RESTful_API/pkg/entities"
	e2 "GO_RESTful_API/pkg/entities/impl/book"
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

func (repo *BookRepository) Create(entity e.Entity) (bool, error) {
	logger.Log("trace", "Book insertion into database begun....")
	book := e2.EntityToBook(entity)
	_, err := repo.DB.Exec(`INSERT INTO "book" VALUES ($1, $2, $3, $4, $5)`,
		repo.uuidGenerator.GenerateUUID(),
		book.Title,
		book.NumberOfPages,
		book.Description,
		book.AuthorID)
	if err != nil {
		logger.Log("errors", err.Error())
		return false, err
	}
	logger.Log("trace", "Book insertion into database finished.")
	return true, err
}

func (repo *BookRepository) Delete(ID string) (bool, error) {
	logger.Log("trace", "Book deleting from database begun....")
	_, err := repo.DB.Exec(`DELETE FROM book WHERE id = $1`, ID)
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false, err
	}
	logger.Log("trace", "Book deleting from database finished.")
	return true, err
}

func (repo *BookRepository) Update(ID string, entity e.Entity) (bool, error) {
	logger.Log("trace", "Book updating from database begun....")
	book := e2.EntityToBook(entity)
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
		logger.Log("errors", err.Error())
		return false, err
	}
	logger.Log("trace", "Book updating from database finished.")
	return false, err
}

func (repo *BookRepository) GetByID(ID string) (e.Entity, error) {
	logger.Log("trace", "Book receiving from database begun....")
	row := repo.DB.QueryRow("SELECT * FROM book WHERE id = $1", ID)
	var book e2.Book
	err := row.Scan(&book.ID, &book.Title, &book.NumberOfPages, &book.Description, &book.AuthorID)
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return nil, err
	}
	logger.Log("trace", "Book receiving from database finished.")
	return &book, err
}

func (repo *BookRepository) GetAll() ([]e.Entity, error) {
	logger.Log("trace", "Receiving all books from database begun....")
	rows, err := repo.DB.Query("SELECT * FROM book")
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return make([]e.Entity, 0), err
	}

	defer rows.Close()

	var res []e.Entity
	for rows.Next() {
		var book e2.Book
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
	return res, err
}

func (repo *BookRepository) Exist(ID string) (bool, error) {
	logger.Log("trace", "Checking whether the book exists....")
	var exists bool
	err := repo.DB.QueryRow("SELECT (exists(SELECT 1 FROM book WHERE id = $1))", ID).Scan(&exists)
	if err != nil {
		logger.Log("errors", err.Error())
		return false, err
	}
	logger.Log("trace", "Checking whether the book exists finished.")
	return exists, err
}

func (repo *BookRepository) Clear() (bool, error) {
	logger.Log("trace", "Clearing books from database begun....")
	_, err := repo.DB.Exec("DELETE FROM book")
	if err != nil {
		log.Printf("%s - [ERROR] %s \n", time.Now(), err.Error())
		return false, err
	}
	logger.Log("trace", "Clearing books from database finished.")
	return true, err
}
