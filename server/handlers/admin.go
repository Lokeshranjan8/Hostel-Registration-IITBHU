package handlers

import (
	"net/http"
	"server/db"
	"server/models"
	"encoding/json"
)


func Getallstudents(w http.ResponseWriter, r *http.Request){

	rows, err := db.DB.Query(`
		SELECT id, student_id, full_name, institute_email, 
		program, branch, current_year, current_semester, is_verified 
		FROM students
	`)

	if err!=nil{
		http.Error(w,"failed to fetch students",http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var students []models.Student
	for rows.Next(){
		var s models.Student
		rows.Scan(
			&s.ID,
			&s.StudentID,
			&s.FullName,
			&s.InstituteEmail,
			&s.Program,
			&s.Branch,
			&s.CurrentYear,
			&s.CurrentSemester,
			&s.IsVerified,
		)
		students = append(students, s)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)

}