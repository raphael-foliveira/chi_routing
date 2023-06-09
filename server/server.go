package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raphael-foliveira/chi_routing/database"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic(err)
	}
	database.RunDb()

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", getMainRouter())
}
