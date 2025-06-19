package services

import "lolymarsh/internal/entity"

type UserService interface {
	RegisterUser() (*entity.UserEntity, error)
}

func (sv *service) UserService() UserService {
	return sv
}

func (sv *service) RegisterUser() (*entity.UserEntity, error) {

	return nil, nil
}
