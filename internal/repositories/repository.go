package repositories

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryAction interface {
	NewUserRepository() UserRepository
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) RepositoryAction {
	return &repository{db: db}
}
