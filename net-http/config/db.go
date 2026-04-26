package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //underline -> import "vazio" do drive do postgres
)

//retorna uma instância do database (postgres nesse projeto)
func SetupDatabase() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil{
		log.Fatal("Error in loading .env file.")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	conn_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	//conexão ao banco
	db, dbErr := sql.Open("postgres", conn_string)
	if dbErr != nil {
		log.Fatal("Error in database connection.")
	}

	//verificação da conexão
	connErr := db.Ping()
	if connErr != nil {
		log.Fatalf("Failed connection to database: %v\nConnection string: %s", connErr, conn_string)
	}
	fmt.Println("Successfully connected to database!")

	return db
}