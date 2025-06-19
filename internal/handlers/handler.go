package handlers

import (
	"lolymarsh/internal/services"
	"lolymarsh/pkg/configs"
)

type Handler struct {
	conf *configs.Config
	sv   *services.ServiceAction
}

func NewHandler(conf *configs.Config, sv *services.ServiceAction) *Handler {
	return &Handler{conf: conf, sv: sv}
}
