package main

import (
	"fmt"
	"log"
	"net/http"
	"server/db"
	"server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Trying to connect to db server ")
	db.Connect()

	router := mux.NewRouter()

	// handlers logic

	router.HandleFunc("admin/students",handlers.Getallstudents ).Methods("GET")




	router.HandleFunc("/", server).Methods("GET")
	fmt.Println("Go rest server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",router))

}

func server(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello from Go-server"))
}