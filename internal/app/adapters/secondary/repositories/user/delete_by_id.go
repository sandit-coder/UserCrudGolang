package repositories

import (
	apperrors "UserCrud/internal/app/domain/errors"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *UserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"

	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete user rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("user: %w", apperrors.ErrNotFound)
	}

	return nil
}
