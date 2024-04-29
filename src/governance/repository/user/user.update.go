package user

import (
	"backend/src/governance/entitiy/user"
	"context"

	"github.com/google/uuid"
)

func (r *RepositoryUser) UpdateExistingUserById(ctx context.Context, id uuid.UUID, user *user.User) (*uuid.UUID, error) {
	db := r.GetDB()

	SQL := `UPDATE users SET name = $2, email = $3, password = $4, avatar = $5 WHERE id = $1;`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, SQL, id.String(), user.Name, user.Email, user.Password, user.Avatar)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
