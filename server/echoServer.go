package server

import (
	"fmt"
	"lolymarsh/pkg/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	conf *configs.Config
	App  *echo.Echo
}

func NewServer(conf *configs.Config) *Server {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.Server.AllowOrigins,
		AllowMethods: conf.Server.AllowMethods,
		AllowHeaders: conf.Server.AllowHeaders,
	}))

	return &Server{conf: conf, App: e}
}

func (s *Server) Run() error {
	conf := s.conf

	fmt.Printf("Starting server at %s\n", conf.Server.PortAPI)

	return s.App.Start(":" + conf.Server.PortAPI)
}
