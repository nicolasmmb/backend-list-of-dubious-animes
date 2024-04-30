package user

import (
	"backend/src/governance/entitiy/user"
	"context"
	"errors"
)

func (r *RepositoryUser) CreateNewUser(ctx context.Context, user *user.User) (*user.User, error) {
	db := r.GetDB()

	SQL := `INSERT INTO users (name, email, password, avatar) VALUES ($1, $2, $3, $4) RETURNING id;`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(ctx, SQL, user.Name, user.Email, user.Password, user.Avatar).Scan(&user.ID)
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
