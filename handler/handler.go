package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/GermanMontejo/thesis-server/model"
	"github.com/GermanMontejo/thesis-server/utils"
	. "github.com/GermanMontejo/thesis-server/repository"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
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

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)
	idStr := ids["id"]
	id, err := strconv.Atoi(idStr)
	log.Println("ids", ids)
	if err != nil {
		log.Println("Error converting string id to integer", err)
		return
	}

	pgsql := utils.InitializeDB()
	repo := Repository{pgsql}
	isDeleted := repo.DeleteStudent(id)
	jsonStr := fmt.Sprint(`{"status_deleted":"`, isDeleted, `"}`)
	w.Write([]byte(jsonStr))
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Println("Error in decoding response body:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("Error: " + err.Error()))
	}
	pgsql := utils.InitializeDB()
	repo := Repository{pgsql}
	isUpdated := repo.UpdateStudent(student)
	jsonStr := fmt.Sprint(`{"status_updated":"`, isUpdated, `"}`)
	w.Write([]byte(jsonStr))
}