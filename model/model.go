package model

import "database/sql"

type Student struct {
	Id		string `json:"userid"`
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	Idnum		string	`json:"idnum"`
	Pass		string	`json:"pass"`
	Schoolyear	string	`json:"schoolyear"`
	Term		string	`json:"term"`
	RoomRequest	string	`json:"roomrequest"`
	Dlsuemail	string	`json:"dlsuemail"`
	Assignroom	string	`json:"assignroom"`
	Status		string	`json:"status"`
	Contactnum	int	`json:"contactnum"`
	MacAddress	string	`json:"mac_address"`
}

type PgSQL struct {
	SQL *sql.DB
}