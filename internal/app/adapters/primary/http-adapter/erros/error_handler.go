package erros

import (
	"UserCrud/internal/app/domain/errors"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, apperrors.ErrInvalidInput):
		return c.Status(400).JSON(errResponse(err))

	case errors.Is(err, apperrors.ErrAlreadyExists):
		return c.Status(409).JSON(errResponse(err))

	case errors.Is(err, apperrors.ErrNotFound):
		return c.Status(404).JSON(errResponse(err))

	case errors.Is(err, context.Canceled),
		errors.Is(err, context.DeadlineExceeded):
		return c.Status(408).JSON(errResponse("timeout"))

	default:
		//	logError(c, err)
		return c.Status(500).JSON(errResponse("internal error"))
	}
}
