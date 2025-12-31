package converters

import (
	"UserCrud/internal/app/application/dtos"
	"UserCrud/internal/app/domain/entities"
)

func ToDtoFromEntity(user *entities.User) *dtos.User {
	return &dtos.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
