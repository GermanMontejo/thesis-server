package utils

import (
	"database/sql"
	. "github.com/GermanMontejo/thesis-server/model"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

var pgsql PgSQL

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "web_server_thesis"
)

func InitializeDB() *PgSQL{

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	if pgsql.SQL == nil {
		db, err := sql.Open("postgres", dbinfo)
		defer db.Close()
		if err != nil {
			log.Fatal("Error while trying to open a db connection:", err)
		}
		log.Println("db stats:", db.Stats())
		pgsql = PgSQL{db}

		_, err = db.Exec(`CREATE TABLE students(id serial, firstname text, lastname text, course text, year int, section string, mac_address string)`)
		if err  != nil {
			log.Fatal("Error while creating a table:", err)
		}
	}
	return &pgsql
}
