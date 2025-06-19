package route

import (
	"lolymarsh/internal/handlers"
	"lolymarsh/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func authRoutes(app *echo.Echo, h *handlers.Handler, m *middlewares.Middleware) {

	app.POST(API_V1+"/auth/register", h.RegisterUser)
	app.POST(API_V1+"/auth/login", h.LoginUser)
}
