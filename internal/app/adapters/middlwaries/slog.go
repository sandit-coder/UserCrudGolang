package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewSlogMiddleware(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqID := uuid.NewString()

		reqLog := log.With(
			"request_id", reqID,
			"method", c.Method(),
			"path", c.Path(),
		)

		c.Locals("logger", reqLog)

		err := c.Next()

		reqLog.Info("request finished",
			"status", c.Response().StatusCode(),
		)

		return err
	}
}
