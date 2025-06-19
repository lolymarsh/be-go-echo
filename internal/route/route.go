package route

import (
	"lolymarsh/internal/handlers"
	"lolymarsh/internal/middlewares"
	"lolymarsh/pkg/configs"

	"github.com/labstack/echo/v4"
)

type Route struct {
	conf *configs.Config
	app  *echo.Echo
	h    *handlers.Handler
	m    *middlewares.Middleware
}

func NewRoute(conf *configs.Config, app *echo.Echo, h *handlers.Handler, m *middlewares.Middleware) *Route {
	return &Route{conf: conf, app: app, h: h, m: m}
}
