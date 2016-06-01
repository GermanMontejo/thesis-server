package repository

import (
	. "github.com/GermanMontejo/thesis-server/model"
	"log"
)

type Repository struct {
	*PgSQL
}

func (repo *Repository) InsertStudent(student *Student) error {
	sql := repo.SQL
	res, err := sql.Exec("INSERT INTO students(firstname, lastname, course, year, section, mac_address) VALUES($1, $2, $3, $4, $5, $6)", student.Firstname, student.Lastname, student.Course, student.Year, student.Section, student.MacAddress)
	if err != nil {
		log.Println("Error inserting data:", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error reading rows affected:", err)
	}
	log.Println("Inserted", rows, "record(s).")
	return nil
}

func (repo *Repository) GetStudents() []Student {
	sql := repo.SQL
	student := new(Student)
	students := []Student{}
	rows, err := sql.Query("SELECT * FROM students")
	if err != nil {
		log.Println("Error retrieving student data:", err)
		return students
	}
	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Firstname, &student.Lastname, &student.Course, &student.Year, &student.Section, &student.MacAddress)
		if err != nil {
			log.Println("Error retrieving data set:", err)
			return []Student{}
		}
		students = append(students, *student)
	}
	log.Println(student)
	return students
}