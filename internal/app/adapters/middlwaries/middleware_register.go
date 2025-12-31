package middleware

import "github.com/gofiber/fiber/v2"

func RegisterMiddleware(app *fiber.App, slogMw fiber.Handler) {
	app.Use(slogMw)
}
