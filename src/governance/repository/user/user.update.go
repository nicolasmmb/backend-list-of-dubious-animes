package user

import (
	"backend/src/governance/entitiy/user"
	"context"
)

func (r *RepositoryUser) UpdateExistingUser(ctx context.Context, user *user.User) (*user.User, error) {
	db := r.GetDB()

	SQL := `UPDATE users SET name = $2, email = $3, password = $4, avatar = $5 WHERE id = $1 RETURNING id;`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(ctx, SQL, user.ID, user.Name, user.Email, user.Password, user.Avatar).Scan(&user.ID, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
