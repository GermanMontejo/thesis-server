package model

import "database/sql"

type Student struct {
	Id		string `json:"_id"`
	Firstname	string	`json:"first_name"`
	Lastname	string	`json:"last_name"`
	Course		string	`json:"course"`
	Year		int	`json:"year"`
	Section		string	`json:"section"`
	MacAddress	string	`json:"mac_address"`
}

type PgSQL struct {
	SQL *sql.DB
}