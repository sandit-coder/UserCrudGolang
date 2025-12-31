package services

import (
	"UserCrud/internal/app/application/dtos"
	"UserCrud/internal/app/domain/entities"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (service *UserService) Create(ctx context.Context, request *dtos.CreateUserRequest) (id uuid.UUID, err error) {

	var entity = entities.NewUser(
		uuid.New(),
		request.Email,
		request.FirstName,
		request.LastName)

	id, err = service.repo.Create(ctx, entity)
	if err != nil {
		fmt.Errorf("soki soki")
	}

	return id, nil
}
