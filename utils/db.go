package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" 
)

var DB *sql.DB

func ConnectDB() {
	var err error

	connStr := "user=your_user password=your_password dbname=your_dbname sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established")
}
