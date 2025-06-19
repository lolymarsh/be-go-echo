package handlers

import (
	"lolymarsh/internal/services"
	"lolymarsh/pkg/configs"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	conf     *configs.Config
	sv       services.ServiceAction
	validate *validator.Validate
}

func NewHandler(conf *configs.Config, sv services.ServiceAction, validate *validator.Validate) *Handler {
	return &Handler{conf: conf, sv: sv, validate: validate}
}
