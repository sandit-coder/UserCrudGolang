package postgres

import (
	"UserCrud/internal/app/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDb(cfg *config.DBConfig) (*sql.DB, error) {
	connectionString := config.GetPostgresConnectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
