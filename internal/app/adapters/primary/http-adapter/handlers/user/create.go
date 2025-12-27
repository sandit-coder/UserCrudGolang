package handlers

import (
	"UserCrud/internal/app/application/dtos"

	"github.com/gofiber/fiber/v2"
)

func (handler *UserHandler) Create(c *fiber.Ctx) error {
	var requestDto dtos.CreateUserRequest

	if err := c.BodyParser(&requestDto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid json")
	}

	id, err := handler.service.Create(c.Context(), &requestDto)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}
