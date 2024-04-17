package main

import (
	"github.com/MarkTBSS/078_Custom_Error/config"
	"github.com/MarkTBSS/078_Custom_Error/databases"
	"github.com/MarkTBSS/078_Custom_Error/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
