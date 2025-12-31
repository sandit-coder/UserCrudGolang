package repositories

import (
	ports "UserCrud/internal/app/application/ports/user"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &UserRepository{
		db: db,
	}
}
