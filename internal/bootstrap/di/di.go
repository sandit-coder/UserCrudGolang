package di

import (
	middleware "UserCrud/internal/app/adapters/middlwaries"
	handlers "UserCrud/internal/app/adapters/primary/http-adapter/handlers/user"
	"UserCrud/internal/app/adapters/primary/http-adapter/router"
	repositories "UserCrud/internal/app/adapters/secondary/repositories/user"
	"UserCrud/internal/app/application/services"
	"UserCrud/internal/app/config"
	server "UserCrud/pkg/fiber"
	"UserCrud/pkg/logger"
	"UserCrud/pkg/postgres"

	"go.uber.org/fx"
)

var ConfigModule = fx.Module("config",
	fx.Provide(
		config.NewFiberConfig,
		config.NewDBConfig),
	fx.Invoke(config.LoadConfig))

var RepositoryModule = fx.Module("repository",
	fx.Provide(
		repositories.NewUserRepository,
	))

var ServiceModule = fx.Module("service",
	fx.Provide(
		services.NewUserService,
	))

var HandlerModule = fx.Module("handler",
	fx.Provide(
		handlers.NewUserHandler))

var PostgresModule = fx.Module("db",
	fx.Provide(
		postgres.NewDb))

var FiberModule = fx.Module("server",
	fx.Provide(
		server.NewFiberApp),
	fx.Invoke(server.RunFiberServer))

var LoggerModule = fx.Module("logger",
	fx.Provide(
		logger.NewLogger,
	))

var MiddlewareModule = fx.Module(
	"middleware",
	fx.Provide(
		middleware.NewSlogMiddleware,
	),
	fx.Invoke(middleware.RegisterMiddleware))

var RouterModule = fx.Module("router",
	fx.Invoke(
		router.AppendUserRoutes))
