package services

import (
	ports "UserCrud/internal/app/application/ports/user"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &UserService{repo}
}
