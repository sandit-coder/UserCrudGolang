package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func LoadConfig() error {
	if err := godotenv.Load("app.env"); err != nil {
		return fmt.Errorf("Error loading .env file", err)
	}
	return nil
}

var Module = fx.Module("config",
	fx.Provide(
		NewFiberConfig,
		NewDBConfig))
