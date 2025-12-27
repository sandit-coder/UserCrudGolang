package repositories

import (
	"UserCrud/internal/app/domain/entities"
	"context"
	"fmt"
)

func (repo *UserRepository) Update(ctx context.Context, entity *entities.User) (err error) {
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"

	_, err = repo.db.ExecContext(ctx, query, entity.FirstName, entity.LastName, entity.Email, entity.ID)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
