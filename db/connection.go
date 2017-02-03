package db

import (
	"database/sql"
	"log"
)

var Connection *sql.DB

func init() {
	con, err := sql.Open("postgres", "user=go dbname=go sslmode=disable")
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	err = con.Ping()
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	log.Println("DB initialized")
	Connection = con
}
