package main

import (
	"github.com/gorilla/mux"
	"github.com/GermanMontejo/thesis-server/handler"
	"net/http"
	"log"
	"time"
)

func main() {

	m := mux.NewRouter().StrictSlash(false)
	m.HandleFunc("/api/students", handlers.PostStudentInfo).Methods("POST")
	m.HandleFunc("/api/students", handlers.GetStudents).Methods("GET")
	m.HandleFunc("/api/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	m.HandleFunc("/api/students", handlers.UpdateStudent).Methods("PUT")
	log.Println("Listening on port:", 8080)
	server := &http.Server{
		Addr: ":8080",
		Handler: m,
		ReadTimeout: time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	server.ListenAndServe()
}
