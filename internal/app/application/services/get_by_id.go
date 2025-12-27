package services

import (
	"UserCrud/internal/app/application/converters"
	"UserCrud/internal/app/application/dtos"
	"context"

	"github.com/google/uuid"
)

func (service *UserService) GetById(ctx context.Context, id uuid.UUID) (*dtos.User, error) {
	user, err := service.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return converters.ToDtoFromEntity(user), err
}
