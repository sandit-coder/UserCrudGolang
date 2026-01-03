package user

import (
	"UserCrud/internal/User/domain/entities"
	"UserCrud/internal/User/usecase/dtos"
	"context"

	"github.com/google/uuid"
)

func (service *Service) Update(ctx context.Context, request *dtos.UpdateUserRequest, id uuid.UUID) error {
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
