package route

import (
	"lolymarsh/internal/handlers"
	"lolymarsh/internal/middlewares"

	"github.com/labstack/echo/v4"
)

const API_V1 = "/api/v1"

func NewRouter(app *echo.Echo, hand *handlers.Handler, midd *middlewares.Middleware) {
	userRoutes(app, hand, midd)
}
