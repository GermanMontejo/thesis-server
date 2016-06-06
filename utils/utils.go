package utils

import (
	"database/sql"
	. "github.com/thesis-server/model"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

var pgsql PgSQL

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "thesis123"
	DB_NAME     = "thesis"
)

func InitializeDB() *PgSQL{

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD,  DB_NAME)

	log.Println("dbinfo:", dbinfo)

	if pgsql.SQL == nil {
		db, err := sql.Open("postgres", dbinfo)
		if err != nil {
			log.Fatal("Error while trying to open a db connection:", err)
		}
		log.Println("db stats:", db.Stats())
		pgsql = PgSQL{db}

		_, err = db.Exec(`CREATE TABLE roomrequest_user2(userid serial, firstname text, lastname text, idnum text, pass text, schoolyear text, term int, roomrequest text, dlsuemail text, assignroom text, status text, contactnum int, mac_address text)`)
		if err  != nil {
			log.Println("Error while creating a table:", err)
		}
	}
	return &pgsql
}
