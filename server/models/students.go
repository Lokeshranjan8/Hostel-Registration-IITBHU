package models

import "time"

type Student struct {
	ID               int       `json:"id"`
	StudentID        string    `json:"student_id"`
	FullName         string    `json:"full_name"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Gender           string    `json:"gender"`
	InstituteEmail   string    `json:"institute_email"`
	PhoneNumber      string    `json:"phone_number"`
	PasswordHash     string    `json:"-"`          
	Program          string    `json:"program"`
	Branch           string    `json:"branch"`
	CurrentYear      int       `json:"current_year"`
	CurrentSemester  int       `json:"current_semester"`
	IsVerified       bool      `json:"is_verified"`
	EnrollmentStatus string    `json:"enrollment_status"`
	CreatedAt        time.Time `json:"created_at"`
}