package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		log.Println("Running without .env (Railway environment)")
	}

	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("fail connect:", err)
	}

	log.Println("Connected to database")

	return db
}
