package converters

import (
	"UserCrud/internal/User/domain/entities"
	"UserCrud/internal/User/usecase/dtos"
)

func ToDtoFromEntity(user *entities.User) *dtos.User {
	return &dtos.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
