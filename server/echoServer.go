package server

import (
	"context"
	"lolymarsh/internal/handlers"
	"lolymarsh/internal/middlewares"
	"lolymarsh/internal/repositories"
	"lolymarsh/internal/route"
	"lolymarsh/internal/services"
	"lolymarsh/pkg/common"
	"lolymarsh/pkg/configs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	conf *configs.Config
	app  *echo.Echo
	db   *sqlx.DB
}

func NewServer(conf *configs.Config, db *sqlx.DB) *Server {

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.Server.AllowOrigins,
		AllowMethods: conf.Server.AllowMethods,
		AllowHeaders: conf.Server.AllowHeaders,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           conf.Server.Format,
		CustomTimeFormat: conf.Server.TimeFormat,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyLimit("20M"))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusNotFound {
			c.JSON(http.StatusNotFound, map[string]string{"error": "wrong path or wrong method"})
			return
		}

		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	vali := common.InitValidate()

	repo := repositories.NewRepository(db)

	serv := services.NewService(conf, db, repo)

	hand := handlers.NewHandler(conf, serv, vali)

	mid := middlewares.NewMiddleware(conf)

	route.NewRouter(e, hand, mid)

	return &Server{conf: conf, app: e, db: db}
}

func (s *Server) Start() error {

	go func() {
		if err := s.app.Start(":" + s.conf.Server.PortAPI); err != nil && err != http.ErrServerClosed {
			s.app.Logger.Fatal("Server start failed: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.db.Close(); err != nil {
		s.app.Logger.Error("Database close failed: ", err)
	}

	return s.app.Shutdown(ctx)
}
