package repositories

import (
	"fmt"
	"lolymarsh/internal/entity"
	"lolymarsh/pkg/common"
)

type UserRepository interface {
	InsertUser(user *entity.UserEntity) error
	GetUserByFilter(filter *common.Filters) (*entity.UserEntity, error)
}

func (r *repository) UserRepository() UserRepository {
	return r
}

func (r *repository) InsertUser(user *entity.UserEntity) error {
	query := `
        INSERT INTO users (
            user_id, first_name, last_name, username, email, password, role, is_active,
            created_at, updated_at, updated_by
        ) VALUES (
            :user_id, :first_name, :last_name, :username, :email, :password, :role, :is_active,
            :created_at, :updated_at, :updated_by
        )`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *repository) GetUserByFilter(filter *common.Filters) (*entity.UserEntity, error) {

	allowedFields := map[string]bool{
		"username": true,
		"email":    true,
		"user_id":  true,
	}
	if !allowedFields[filter.Field] {
		return nil, fmt.Errorf("invalid filter field: %s", filter.Field)
	}

	query := `
        SELECT user_id, first_name, last_name, username, email, password, role, is_active,
            created_at, updated_at, updated_by
        FROM users
        WHERE %s = ?
    `

	query = fmt.Sprintf(query, filter.Field)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(filter.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var user entity.UserEntity
	if rows.Next() {
		if err := rows.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.UpdatedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
	} else {
		return nil, nil
	}

	return &user, nil
}
