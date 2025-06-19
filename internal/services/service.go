package services

import (
	"github.com/jmoiron/sqlx"

	"lolymarsh/internal/repositories"
	"lolymarsh/pkg/configs"
)

type ServiceAction interface {
	UserService() UserService
}

type service struct {
	conf *configs.Config
	db   *sqlx.DB
	repo *repositories.RepositoryAction
}

func NewService(conf *configs.Config, db *sqlx.DB, repo *repositories.RepositoryAction) ServiceAction {
	return &service{conf: conf, db: db, repo: repo}
}
