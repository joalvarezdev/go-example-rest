package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := "user=postgres password=postgres dbname=go-gpt sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	log.Println("Connected to PostgreSQL database")
}