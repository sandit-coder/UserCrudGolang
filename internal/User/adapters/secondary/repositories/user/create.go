package repositories

import (
	"context"
	"errors"
	"fmt"

	"UserCrud/internal/User/domain/entities"
	"UserCrud/internal/User/domain/errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *UserRepository) Create(ctx context.Context, entity *entities.User) (uuid.UUID, error) {
	query := `
        INSERT INTO users (id, email, first_name, last_name)
        VALUES ($1, $2, $3, $4)
    `

	_, err := repo.db.Exec(ctx, query, entity.ID, entity.Email, entity.FirstName, entity.LastName)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return uuid.Nil, fmt.Errorf("user: %w", apperrors.ErrAlreadyExists)
			}
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return uuid.Nil, err
		}

		return uuid.Nil, fmt.Errorf("repo create user: %w", err)
	}

	return entity.ID, nil
}
