package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "159357"
	dbname   = "books"
)

var DB *sql.DB

func init() {
	var err error

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}
