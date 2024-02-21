package book

import (
	"GO_RESTful_API/pkg/entities"
	"log"
	"time"
)

type Book struct {
	ID            string
	Title         string
	NumberOfPages int
	Description   string
	AuthorID      string
}

func (b *Book) GetID() string {
	return b.ID
}

func EntityToBook(e entities.Entity) *Book {
	book, ok := e.(*Book)
	if !ok {
		log.Printf("%s - [ERROR] Passed entities is not of type Book", time.Now())
		return nil
	}
	return book
}
