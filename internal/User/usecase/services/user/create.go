package user

import (
	"UserCrud/internal/User/domain/entities"
	"UserCrud/internal/User/usecase/dtos"
	"context"

	"github.com/google/uuid"
)

func (service *Service) Create(ctx context.Context, request *dtos.CreateUserRequest) (id uuid.UUID, err error) {

	var entity = entities.NewUser(
		uuid.New(),
		request.Email,
		request.FirstName,
		request.LastName)

	id, err = service.repo.Create(ctx, entity)
	if err != nil {
		return id, err
	}

	return id, nil
}
