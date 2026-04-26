package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

func SetupDatabase() *gorm.DB {
	once.Do(func () {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error in enviroment variables.")
		}

		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbPort := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort,
		)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database.")
		}

		fmt.Println("Successfully connected to database.")
	})

	return db
}