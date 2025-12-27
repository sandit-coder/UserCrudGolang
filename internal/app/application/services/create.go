package services

import (
	"UserCrud/internal/app/application/converters"
	"UserCrud/internal/app/application/dtos"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (service *UserService) Create(ctx context.Context, req *dtos.CreateUserRequest) (id uuid.UUID, err error) {
	var dto = converters.ToDtoFromCreateUserRequest(req)
	dto.ID = uuid.New()

	id, err = service.repo.Create(ctx, converters.ToEntityFromDto(dto))
	if err != nil {
		fmt.Errorf("soki soki")
	}

	return id, nil
}
