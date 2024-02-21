package book

import (
	"GO_RESTful_API/pkg/entities"
	"log"
	"time"
)

type Author struct {
	ID        string
	Name      string
	Surname   string
	Birthdate time.Time
	DeathDate time.Time
}

func (a *Author) GetID() string {
	return a.ID
}
func EntityToAuthor(e entities.Entity) *Author {
	author, ok := e.(*Author)
	if !ok {
		log.Printf("%s - [ERROR] Passed entities is not of type Author", time.Now())
		return nil
	}
	return author
}
