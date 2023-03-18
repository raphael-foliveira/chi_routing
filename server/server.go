package server

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raphael-foliveira/hot_reload/database"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic(err)
	}
	database.RunDb()

	http.ListenAndServe(":8000", getMainRouter())
}
