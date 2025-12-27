package handlers

import (
	"UserCrud/internal/app/application/ports/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService, app *fiber.App) *UserHandler {

	return &UserHandler{service: service}
}
