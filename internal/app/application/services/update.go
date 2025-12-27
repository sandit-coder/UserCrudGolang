package services

import (
	"UserCrud/internal/app/application/converters"
	"UserCrud/internal/app/application/dtos"
	"context"

	"github.com/google/uuid"
)

func (service *UserService) Update(ctx context.Context, request *dtos.UpdateUserRequest, id uuid.UUID) error {
	_, err := service.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	var dto = converters.ToDtoFromUpdateUserRequest(request, id)

	if err := service.repo.Update(ctx, converters.ToEntityFromDto(dto)); err != nil {
		return err
	}

	return nil
}
