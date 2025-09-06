package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() *sql.DB {
	var err error
	DB, err = sql.Open("sqlite3", "internal/db/student.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to SQLite database!")
	return DB
}
