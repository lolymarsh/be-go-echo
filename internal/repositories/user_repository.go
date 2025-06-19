package repositories

import (
	"fmt"
	"lolymarsh/internal/entity"
	"lolymarsh/pkg/common"
	"lolymarsh/pkg/util"
	"strings"
)

type UserRepository interface {
	InsertUser(user *entity.UserEntity) error
	GetUserByFilter(filter *common.Filters) (*entity.UserEntity, error)
	FilterUser(filter *common.FilterRequest) ([]*entity.UserEntity, int64, error)
}

func (r *Repository) UserRepository() UserRepository {
	return r
}

func (r *Repository) InsertUser(user *entity.UserEntity) error {
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

func (r *Repository) GetUserByFilter(filter *common.Filters) (*entity.UserEntity, error) {

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

func (r *Repository) FilterUser(filter *common.FilterRequest) ([]*entity.UserEntity, int64, error) {
	conditions := []string{}
	params := []any{}

	for _, f := range filter.Filters {
		switch f.Field {
		// case "search":
		// 	if util.StringIsNotEmpty(f.Value) {
		// 		conditions = append(conditions, fmt.Sprintf("(first_name LIKE ? OR last_name LIKE ? OR username LIKE ? OR email LIKE ?)", "%"+f.Value+"%", "%"+f.Value+"%", "%"+f.Value+"%", "%"+f.Value+"%"))
		// 		params = append(params, "%"+f.Value+"%")
		// 	}
		case "first_name", "last_name", "username", "email", "role":
			if util.StringIsNotEmpty(f.Value) {
				conditions = append(conditions, fmt.Sprintf("%s LIKE ?", f.Field))
				params = append(params, "%"+f.Value+"%")
			}
		case "created_at", "updated_at":
			if f.GreaterThan > 0 {
				conditions = append(conditions, fmt.Sprintf("%s >= ?", f.Field))
				params = append(params, f.GreaterThan)
			}
			if f.LessThan > 0 {
				conditions = append(conditions, fmt.Sprintf("%s <= ?", f.Field))
				params = append(params, f.LessThan)
			}
		case "is_active":
			if util.StringIsNotEmpty(f.Value) && f.Value == "true" || f.Value == "false" {
				conditions = append(conditions, fmt.Sprintf("%s = ?", f.Field))
				params = append(params, f.Value == "true")
			}
		default:
			continue
		}
	}

	whereClause := "1=1"
	if len(conditions) > 0 {
		whereClause = strings.Join(conditions, " AND ")
	}

	query := `
		SELECT user_id, first_name, last_name, username, email, password, role, is_active,
			created_at, updated_at, updated_by
		FROM users
		WHERE %s
	`

	sortClause := ""
	if filter.SortName != "" && filter.SortBy != "" {
		validSortFields := map[string]bool{
			"user_id":    true,
			"first_name": true,
			"last_name":  true,
			"username":   true,
			"email":      true,
			"role":       true,
			"is_active":  true,
			"created_at": true,
			"updated_at": true,
		}
		if validSortFields[filter.SortName] && util.StringIsNotEmpty(filter.SortBy) {
			sortClause = fmt.Sprintf("ORDER BY %s %s", filter.SortName, filter.SortBy)
		}
	}

	limit := filter.PageSize
	offset := (filter.Page - 1) * filter.PageSize
	query = fmt.Sprintf("%s %s LIMIT %d OFFSET %d", fmt.Sprintf(query, whereClause), sortClause, limit, offset)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	users := []*entity.UserEntity{}
	for rows.Next() {
		var user entity.UserEntity
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
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, &user)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(user_id) FROM users WHERE %s", whereClause)
	var total int64
	err = r.db.QueryRow(countQuery, params...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %w", err)
	}

	return users, total, nil
}
