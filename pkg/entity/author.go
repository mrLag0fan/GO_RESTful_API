package entity

import (
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
func EntityToAuthor(e Entity) *Author {
	author, ok := e.(*Author)
	if !ok {
		log.Printf("%s - [ERROR] Passed entity is not of type Author", time.Now())
		return nil
	}
	return author
}
