package server

import (
	"UserCrud/internal/app/config"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewFiberApp(cfg *config.FiberConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	return app
}

func RunFiberServer(
	lc fx.Lifecycle,
	app *fiber.App,
	cfg *config.FiberConfig,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Printf("Starting HTTP server on %s", cfg.Port)
				if err := app.Listen(cfg.Port); err != nil {
					log.Printf("Fiber server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down Fiber server...")
			return app.Shutdown()
		},
	})
}
