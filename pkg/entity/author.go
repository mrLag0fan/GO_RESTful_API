package entity

import "time"

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
