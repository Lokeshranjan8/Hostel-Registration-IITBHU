package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/auth"
	"server/db"
	"server/models"
)

func RegisterStudent(w http.ResponseWriter, r *http.Request) {
	var body struct {
		StudentID       string `json:"student_id"`
		FullName        string `json:"full_name"`
		DateOfBirth     string `json:"date_of_birth"`
		Gender          string `json:"gender"`
		InstituteEmail  string `json:"institute_email"`
		PhoneNumber     string `json:"phone_number"`
		Password        string `json:"password"`
		Program         string `json:"program"`
		Branch          string `json:"branch"`
		CurrentYear     int    `json:"current_year"`
		CurrentSemester int    `json:"current_semester"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.StudentID == "" || body.FullName == "" || body.Password == "" ||
		body.InstituteEmail == "" || body.Program == "" || body.Branch == "" {
		http.Error(w, "Required fields missing", http.StatusBadRequest)
		return
	}

	hash, err := auth.HashPassword(body.Password)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec(`
		INSERT INTO students 
		(student_id, full_name, date_of_birth, gender, institute_email, phone_number, password_hash, program, branch, current_year, current_semester)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		body.StudentID, body.FullName, body.DateOfBirth, body.Gender,
		body.InstituteEmail, body.PhoneNumber, hash,
		body.Program, body.Branch, body.CurrentYear, body.CurrentSemester,
	)
	if err != nil {
		http.Error(w, "Could not register student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Student registered successfully",
	})
}

func LoginStudent(w http.ResponseWriter, r *http.Request) {
	var body struct {
		StudentID string `json:"student_id"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.StudentID == "" || body.Password == "" {
		http.Error(w, "Student ID and password required", http.StatusBadRequest)
		return
	}

	var student models.Student
	err := db.DB.QueryRowx(
		"SELECT id, password_hash FROM students WHERE student_id = $1",
		body.StudentID,
	).Scan(&student.ID, &student.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if !auth.CheckPassword(student.PasswordHash, body.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(student.ID, "student")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
