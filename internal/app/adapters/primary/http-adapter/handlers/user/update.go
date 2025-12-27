package handlers

import (
	"UserCrud/internal/app/application/dtos"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (handler *UserHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be a valid UUID")
	}

	var req dtos.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid json")
	}

	if err := handler.service.Update(c.Context(), &req, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
