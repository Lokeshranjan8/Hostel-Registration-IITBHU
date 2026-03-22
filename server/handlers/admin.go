package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/auth"
	"server/db"
	"server/models"
)

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.Email == "" || body.Password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	var admin models.Admin
	err := db.DB.QueryRowx(
		"SELECT id, password_hash FROM admins WHERE email = $1",
		body.Email,
	).Scan(&admin.ID, &admin.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if !auth.CheckPassword(admin.PasswordHash, body.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(admin.ID, "admin")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT id, student_id, full_name, institute_email,
		program, branch, current_year, current_semester, is_verified
		FROM students
	`)
	if err != nil {
		http.Error(w, "Failed to fetch students", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(
			&s.ID, &s.StudentID, &s.FullName,
			&s.InstituteEmail, &s.Program, &s.Branch,
			&s.CurrentYear, &s.CurrentSemester, &s.IsVerified,
		)
		students = append(students, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
