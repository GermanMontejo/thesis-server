package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/thesis-server/model"
	"github.com/thesis-server/utils"
	. "github.com/thesis-server/repository"
)

var student model.Student

func PostStudentInfo(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		log.Println("Error in decoding response body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("Error: " + err.Error()))
	}

	pgsql := utils.InitializeDB()
	repo := Repository{pgsql}
	repo.InsertStudent(&student)
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	pgsql := utils.InitializeDB()
	repo := Repository{pgsql}
	students := repo.GetStudents()
	log.Println("Students:", students)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(&students)
	if err != nil {
		log.Println("Error in marshaling students to json format.")
		return
	}
	w.Write(j)
}
