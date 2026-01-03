package user

import (
	"UserCrud/internal/User/usecase/ports/user"
)

type Service struct {
	repo ports.UserRepository
}

func NewService(repo ports.UserRepository) ports.UserService {
	return &Service{repo}
}
