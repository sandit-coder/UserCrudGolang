package repositories

import (
	"UserCrud/internal/app/domain/entities"
	"UserCrud/internal/app/domain/errors"
	"context"
	"errors"
	"fmt"
)

func (repo *UserRepository) Update(ctx context.Context, entity *entities.User) (err error) {
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"

	result, err := repo.db.ExecContext(ctx, query, entity.FirstName, entity.LastName, entity.Email, entity.ID)

	if err != nil {
		if errors.Is(err, context.Canceled) ||
			errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("update user rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("user: %w", apperrors.ErrNotFound)
	}

	return nil
}
