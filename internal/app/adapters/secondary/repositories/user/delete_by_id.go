package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *UserRepository) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	query := "DELETE FROM users WHERE id = $1;"

	_, err = repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
