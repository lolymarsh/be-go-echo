package main

import (
	"lolymarsh/pkg/configs"
	"lolymarsh/pkg/database"
	"lolymarsh/server"

	"github.com/labstack/gommon/log"
)

func main() {
	conf := configs.LoadConfig()

	db, err := database.InitDatabaseMySQL(conf.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	srv := server.NewServer(conf, db)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server start failed: %v", err)
	}
}
