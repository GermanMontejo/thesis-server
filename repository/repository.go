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

func (repo *Repository) GetStudent(id string) *Student {
	sql := repo.SQL
	student := new(Student)
	row := sql.QueryRow("SELECT * FROM students WHERE id = $1", id)
	row.Scan(student.Id, student.Firstname, student.Lastname, student.Course, student.Year, student.Section, student.MacAddress)
	return student
}

func (repo *Repository) DeleteStudent(id int) bool{
	sql := repo.SQL
	res, err := sql.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		log.Println("Error while deleting student data:", err)
		return false
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("There was a problem while trying to delete the row with the passed id:", err.Error())
		return false
	}
	log.Println("There where ", rowsAffected, " row(s) deleted...")
	return true
}

func (repo *Repository) UpdateStudent(student Student) bool {
	sql := repo.SQL
	res, err := sql.Exec("UPDATE students SET firstname=$1, lastname=$2, course=$3, year=$4, section=$5, mac_address=$6 WHERE id=$7", student.Firstname, student.Lastname, student.Course, student.Year, student.Section, student.MacAddress, student.Id)
	if err != nil {
		log.Println("There was a problem while trying to update the row with the passed id:", err.Error())
		return false
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("There was a problem while trying to update the row with the passed id:", err.Error())
		return false
	}
	log.Println("There where ", rowsAffected, " row(s) updated...")
	return true
}