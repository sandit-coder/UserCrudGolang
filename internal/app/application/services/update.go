package services

import (
	"UserCrud/internal/app/application/dtos"
	"UserCrud/internal/app/domain/entities"
	"context"

	"github.com/google/uuid"
)

func (service *UserService) Update(ctx context.Context, request *dtos.UpdateUserRequest, id uuid.UUID) error {
	_, err := service.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	var entity = entities.NewUser(
		uuid.New(),
		request.Email,
		request.FirstName,
		request.LastName)

	if err := service.repo.Update(ctx, entity); err != nil {
		return err
	}

	return nil
}
