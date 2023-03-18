package main

import (
	"github.com/raphael-foliveira/hot_reload/database"
	"github.com/raphael-foliveira/hot_reload/server"
)

func main() {
	server.Run()
	defer database.Db.Close()
}
