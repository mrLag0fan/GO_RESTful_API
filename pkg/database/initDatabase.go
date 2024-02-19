package database

import (
	"GO_RESTful_API/pkg/logger"
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
	logger.Log("trace", "Initializing new database connection")
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
	logger.Log("info", "Database connected!")
	logger.Log("trace", "Initializing new database connection")
}
