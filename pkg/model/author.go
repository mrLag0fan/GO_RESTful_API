package model

import "time"

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birthdate time.Time `json:"birth-date"`
	DeathDate time.Time `json:"death-date"`
}
