package repositories

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryAction interface {
	UserRepository() UserRepository
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) RepositoryAction {
	return &Repository{db: db}
}
