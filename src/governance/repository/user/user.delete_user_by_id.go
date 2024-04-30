package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (r *RepositoryUser) DeleteUserByID(ctx context.Context, id uuid.UUID) (*uuid.UUID, error) {
	db := r.GetDB()

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	// SQL := `UPDATE users SET deleted_at = NOW() WHERE id = $1;`
	SQL := `UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;`

	_, err = tx.Exec(ctx, SQL, id)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	switch err {
	case nil:
		return &id, nil
	case context.Canceled:
		return nil, errors.New("--> Context canceled")
	default:
		return nil, err
	}
}
