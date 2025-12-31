package repositories

import (
	"UserCrud/internal/app/domain/entities"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (repo *UserRepository) Create(ctx context.Context, entity *entities.User) (id uuid.UUID, err error) {
	query := "INSERT INTO users VALUES($1, $2, $3, $4)"

	_, err = repo.db.ExecContext(ctx, query, entity.ID, entity.FirstName, entity.LastName, entity.Email)

	if err != nil {
		if errors.Is(err, context.Canceled) ||
			errors.Is(err, context.DeadlineExceeded) {
			return uuid.Nil, err
		}
		return uuid.Nil, fmt.Errorf("repo create user: %w", err)
	}

	return entity.ID, nil
}
