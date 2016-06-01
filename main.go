package main

import (
	"github.com/gorilla/mux"
	"github.com/GermanMontejo/thesis-server/handler"
	"net/http"
	"log"
)

func main() {

	m := mux.NewRouter().StrictSlash(false)
	m.HandleFunc("/api/students", handlers.PostStudentInfo).Methods("POST")
	m.HandleFunc("/api/students", handlers.GetStudents).Methods("GET")
	log.Println("Listening on port:", 8080)
	http.ListenAndServe(":8080", m)
}
