package erros

import "github.com/gofiber/fiber/v2"

func errResponse(err any) fiber.Map {
	return fiber.Map{
		"error": err,
	}
}
