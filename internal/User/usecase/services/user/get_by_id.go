package user

import (
	"UserCrud/internal/User/usecase/converters"
	"UserCrud/internal/User/usecase/dtos"
	"context"

	"github.com/google/uuid"
)

func (service *Service) GetById(ctx context.Context, id uuid.UUID) (*dtos.User, error) {
	user, err := service.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return converters.ToDtoFromEntity(user), err
}
