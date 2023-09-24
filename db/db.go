package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "articlesdb"
)

func Connect() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to db")
		panic(err)
	}
	log.Println("Successfully connected to db!")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}
