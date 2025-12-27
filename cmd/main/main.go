// cmd/restaurant/main.go
package main

import (
	"UserCrud/internal/app/adapters/primary/http-adapter/handlers/user"
	"UserCrud/internal/app/adapters/primary/http-adapter/router"
	"UserCrud/internal/app/adapters/secondary/repositories/user"
	"UserCrud/internal/app/application/ports/user"
	"UserCrud/internal/app/application/services"
	"UserCrud/internal/app/config"
	"UserCrud/internal/pkg/fiber"
	"UserCrud/internal/pkg/postgres"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		fx.Provide(
			postgres.NewDb,
			fx.Annotate(
				repositories.NewUserRepository,
				fx.As(new(ports.UserRepository)),
			),
		),

		fx.Provide(
			services.NewUserService,
		),

		fx.Provide(
			handlers.NewUserHandler,
		),

		fx.Provide(
			server.NewFiberApp,
		),

		fx.Invoke(config.LoadConfig),
		fx.Invoke(router.AppendUserRoutes),
		fx.Invoke(server.RunFiberServer),
	).Run()
}
