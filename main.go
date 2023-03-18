package main

import (
	"github.com/raphael-foliveira/chi_routing/database"
	"github.com/raphael-foliveira/chi_routing/server"
)

func main() {
	server.Run()
	defer database.Db.Close()
}
