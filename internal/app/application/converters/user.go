package converters

import (
	"UserCrud/internal/app/application/dtos"
	"UserCrud/internal/app/domain/entities"

	"github.com/google/uuid"
)

func ToEntityFromDto(user *dtos.User) *entities.User {

	return &entities.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ToDtoFromEntity(user *entities.User) *dtos.User {

	return &dtos.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ToDtoFromCreateUserRequest(req *dtos.CreateUserRequest) *dtos.User {
	return &dtos.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
}

func ToDtoFromUpdateUserRequest(req *dtos.UpdateUserRequest, id uuid.UUID) *dtos.User {
	return &dtos.User{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
}
