package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/GermanMontejo/thesis-server/model"
	"github.com/GermanMontejo/thesis-server/utils"
	. "github.com/GermanMontejo/thesis-server/repository"
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
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	pgsql := utils.InitializeDB()
	repo := Repository{pgsql}
	students := repo.GetStudents()
	log.Println("Students:", students)

}