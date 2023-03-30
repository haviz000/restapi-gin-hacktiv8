package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "jaka"
	dbname   = "golang_book"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	log.Println("successfully connected to database")

	return db
}
