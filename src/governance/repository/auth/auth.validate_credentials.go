package auth

import (
	entity "backend/src/governance/entity/user"
	"context"
	"errors"
)

func (r *RepositoryAuth) GetUserToValidateCredentials(ctx context.Context, email string) (*entity.User, error) {
	db := r.GetDB()

	SQL := `SELECT id, name, email, password, avatar, created_at, updated_at, deleted_at FROM users WHERE email ilike $1 LIMIT 1;`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	user := &entity.User{}
	err = tx.QueryRow(ctx, SQL, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	switch err {
	case nil:
		return user, nil
	case context.Canceled:
		tx.Rollback(ctx)
		return nil, errors.New("--> Context canceled")
	default:
		tx.Rollback(ctx)
		return nil, err
	}
}
