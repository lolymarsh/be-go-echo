package repositories

import (
	"fmt"
	"lolymarsh/internal/entity"
	"lolymarsh/pkg/common"
)

type UserRepository interface {
	InsertUser(user *entity.UserEntity) error
}

func (r *repository) NewUserRepository() UserRepository {
	return r
}

func (r *repository) InsertUser(user *entity.UserEntity) error {
	query := `
        INSERT INTO users (
            user_id, first_name, last_name, username, email, password,
            created_at, updated_at, updated_by
        ) VALUES (
            :user_id, :first_name, :last_name, :username, :email, :password,
            :created_at, :updated_at, :updated_by
        )`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *repository) GetUserByFilter(filter *common.Filters) (*entity.UserEntity, error) {

	query := `
		SELECT user_id, first_name, last_name, username, email, password,
			created_at, updated_at, updated_by
		FROM users
		WHERE :field = :value
	`

	filterMap := map[string]any{
		"field": filter.Field,
		"value": filter.Value,
	}

	var user *entity.UserEntity
	err := r.db.Get(&user, query, filterMap)
	if err != nil {
		return nil, err
	}

	return user, nil
}
