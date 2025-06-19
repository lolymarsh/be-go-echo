package common

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type FilterRequest struct {
	Filters  []*Filters `json:"keywords" validate:"dive"`
	SortName string     `json:"sort_name" validate:"omitempty"`
	SortBy   string     `json:"sort_by" validate:"omitempty,oneof=asc desc"`
	Page     int        `json:"page" validate:"required,gt=0"`
	PageSize int        `json:"page_size" validate:"required,gt=0,lte=100"`
}

type Filters struct {
	Field       string `json:"field" validate:"required"`
	Value       string `json:"value"`
	GreaterThan int64  `json:"greater_than" validate:"omitempty"`
	LessThan    int64  `json:"less_than" validate:"omitempty"`
}

func InitValidate() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("trim", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		trimmed := strings.Join(strings.Fields(field), "")
		fl.Field().SetString(trimmed)
		return true
	})

	return validate
}
