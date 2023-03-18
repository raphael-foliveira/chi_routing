package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func getEnv() []any {
	dbHost, found := os.LookupEnv("DB_HOST")
	if !found {
		dbHost = "localhost"
	}
	dbPort, found := os.LookupEnv("DB_PORT")
	if !found {
		dbPort = "5432"
	}
	dbUser, found := os.LookupEnv("DB_USER")
	if !found {
		dbUser = "postgres"
	}
	dbPassword, found := os.LookupEnv("DB_PASSWORD")
	if !found {
		dbPassword = "123"
	}
	dbName, found := os.LookupEnv("DB_NAME")
	if !found {
		dbName = "golang_server"
	}
	return []any{dbUser, dbPassword, dbHost, dbPort, dbName}
}

var Db *sql.DB
var err error

func RunDb() {
	Db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", getEnv()...))
	if err != nil {
		log.Println("Error connecting to database")
		panic(err)
	}

	Db.Exec(`set search_path='golang_server'`)

}
