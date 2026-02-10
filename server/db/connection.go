package db

import (
	"log"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

)

var DB *sqlx.DB

func Connect(){
	log.Println("Entering connection go file")

	dsn := os.Getenv("Database_url")
	// Database_url="postgres://postgres:postgres@postgres:5432/hostel-db?sslmode=disable"

	var err error
	DB,err := sqlx.Connect("postgres",dsn)
	if err!=nil {
		log.Fatalln("DB connection failed",err)
	}

	if err:= DB.Ping(); err!=nil{
		log.Fatal("DB ping failed",err)
	}

	log.Println("Successfully connected to database")
}