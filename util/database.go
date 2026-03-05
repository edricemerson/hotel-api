package util

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in env file")
	}

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Not reachable:", err)
	}

	log.Println("Connected successfully")

	DB = db
}
