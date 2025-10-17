package handlers

import (
	"errors"
	"lolymarsh/internal/request"
	"lolymarsh/pkg/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Handler) RegisterUser(ctx echo.Context) error {
	request := &request.RegisterRequest{}

	if err := ctx.Bind(request); err != nil {
		return common.HandleError(ctx, errors.New("failed to parse request body"), http.StatusBadRequest)
	}

	if err := c.validate.Struct(request); err != nil {
		return common.HandleError(ctx, err, http.StatusBadRequest)
	}

	dataService, err := c.sv.UserService().RegisterUser(request)
	if err != nil {
		return common.HandleError(ctx, err)
	}

	return common.HandleSuccess(ctx, http.StatusCreated, "register success", echo.Map{
		"data": dataService,
	})
}

func (c *Handler) LoginUser(ctx echo.Context) error {
	request := &request.LoginRequest{}

	if err := ctx.Bind(request); err != nil {
		return common.HandleError(ctx, errors.New("failed to parse request body"), http.StatusBadRequest)
	}

	if err := c.validate.Struct(request); err != nil {
		return common.HandleError(ctx, err, http.StatusBadRequest)
	}

	dataService, authToken, err := c.sv.UserService().LoginUser(request)
	if err != nil {
		return common.HandleError(ctx, err)
	}

	return common.HandleSuccess(ctx, http.StatusCreated, "login success", echo.Map{
		"data":       dataService,
		"auth_token": authToken,
	})
}

func (c *Handler) FilterUser(ctx echo.Context) error {
	request := &common.FilterRequest{}

	if err := ctx.Bind(request); err != nil {
		return common.HandleError(ctx, errors.New("failed to parse request body"), http.StatusBadRequest)
	}

	if err := c.validate.Struct(request); err != nil {
		return common.HandleError(ctx, err, http.StatusBadRequest)
	}

	dataService, totalDataService, err := c.sv.UserService().FilterUser(request)
	if err != nil {
		return common.HandleError(ctx, err)
	}

	return common.HandleSuccess(ctx, http.StatusCreated, "login success", echo.Map{
		"data":       dataService,
		"total_data": totalDataService,
	})
}
