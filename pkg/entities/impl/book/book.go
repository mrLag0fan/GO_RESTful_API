package book

import (
	"GO_RESTful_API/pkg/entities"
	"log"
	"time"
)

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	NumberOfPages int    `json:"number-of-pages"`
	Description   string `json:"description"`
	AuthorID      string `json:"author-id"`
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
