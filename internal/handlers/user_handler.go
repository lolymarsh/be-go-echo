package handlers

import (
	"errors"
	"lolymarsh/internal/request"
	"lolymarsh/pkg/common"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

func (c *Handler) RegisterUser(ctx echo.Context) error {
	request := &request.RegisterRequest{}

	if err := ctx.Bind(request); err != nil {
		return common.HandleError(ctx, errors.New("failed to parse request body"), fiber.StatusBadRequest)
	}

	if err := c.validate.Struct(request); err != nil {
		return common.HandleError(ctx, err, fiber.StatusBadRequest)
	}

	dataService, err := c.sv.UserService().RegisterUser(request)
	if err != nil {
		return common.HandleError(ctx, err)
	}

	return common.HandleSuccess(ctx, fiber.StatusCreated, "register success", fiber.Map{
		"data": dataService,
	})
}

func (c *Handler) LoginUser(ctx echo.Context) error {
	request := &request.LoginRequest{}

	if err := ctx.Bind(request); err != nil {
		return common.HandleError(ctx, errors.New("failed to parse request body"), fiber.StatusBadRequest)
	}

	if err := c.validate.Struct(request); err != nil {
		return common.HandleError(ctx, err, fiber.StatusBadRequest)
	}

	dataService, authToken, err := c.sv.UserService().LoginUser(request)
	if err != nil {
		return common.HandleError(ctx, err)
	}

	return common.HandleSuccess(ctx, fiber.StatusCreated, "register success", fiber.Map{
		"data":       dataService,
		"auth_token": authToken,
	})
}
