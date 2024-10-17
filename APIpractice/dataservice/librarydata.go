package dataservice

import (
	"database/sql"
	"net/http"
	"my_module/model"
	"encoding/json"
)

func CreateStudent(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
         var student model.Student
		 if err := json.NewDecoder(r.Body).Decode(&student) ; err != nil {
			return err
		}
       query := `INSERT INTO students(id, firstname , lastname, city) VALUES(?,?,?,?)`
		_,err := db.Exec(query,student.Id,student.FirstName,student.LastName,student.City)

		if err != nil {
			return err
		}
		w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(student)
		return nil


}