package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", server).Methods("GET")
	fmt.Println("Go rest server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",router))

}

func server(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello from Go-server"))
}