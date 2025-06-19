package main

import (
	"lolymarsh/internal/handlers"
	"lolymarsh/internal/middlewares"
	"lolymarsh/internal/repositories"
	"lolymarsh/internal/route"
	"lolymarsh/internal/services"
	"lolymarsh/pkg/configs"
	"lolymarsh/pkg/database"
	"lolymarsh/server"

	"github.com/labstack/gommon/log"
)

func main() {
	conf := configs.LoadConfig()

	db, err := database.InitDatabaseSQLite(conf.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server := server.NewServer(conf)

	repo := repositories.NewRepository(db)

	serv := services.NewService(conf, db, &repo)

	hand := handlers.NewHandler(conf, &serv)

	mid := middlewares.NewMiddleware(conf)

	route.NewRoute(conf, server.App, hand, mid)

	server.Run()
}
