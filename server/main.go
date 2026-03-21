package main

import (
	"fmt"
	"log"
	"net/http"
	"server/auth"
	"server/db"
	"server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Trying to connect to db server")
	db.Connect()

	router := mux.NewRouter()

	// public routes
	router.HandleFunc("/api/register", handlers.RegisterStudent).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginStudent).Methods("POST")
	router.HandleFunc("/api/admin/login", handlers.LoginAdmin).Methods("POST")

	// protected routes
	router.HandleFunc("/api/admin/students", auth.RequireAdmin(handlers.GetAllStudents)).Methods("GET")

	fmt.Println("Go rest server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
