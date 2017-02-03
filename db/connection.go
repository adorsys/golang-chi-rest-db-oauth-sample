package db

import (
	"database/sql"
	"github.com/adorsys/golang-chi-rest-db-oauth-sample/config"
	"log"
)

var Connection *sql.DB

func Configure() {
	con, err := sql.Open("postgres", config.Conf.Db.Url)
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
