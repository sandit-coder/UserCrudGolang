package repositories

import (
	"UserCrud/internal/User/domain/entities"
	apperrors "UserCrud/internal/User/domain/errors"
	"context"
	"errors"
	"fmt"
)

func (repo *UserRepository) Update(ctx context.Context, entity *entities.User) (err error) {
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"

	commandTag, err := repo.db.Exec(ctx, query, entity.FirstName, entity.LastName, entity.Email, entity.ID)

	if err != nil {
		if errors.Is(err, context.Canceled) ||
			errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user: %w", apperrors.ErrNotFound)
	}

	return nil
}
