package postgres

import (
	"UserCrud/internal/User/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDb(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	connectionString := config.GetPostgresConnectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
