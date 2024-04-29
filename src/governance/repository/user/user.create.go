package user

import (
	"backend/src/governance/entitiy/user"
	"context"
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
	if err != nil {
		return nil, err
	}

	return user, nil
}
