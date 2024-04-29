package user

import (
	"backend/src/governance/entitiy/user"
	"context"

	"github.com/google/uuid"
)

func (r *RepositoryUser) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	db := r.GetDB()

	SQL := `SELECT id, name, email, password, avatar, created_at, updated_at, deleted_at FROM users WHERE id = $1`

	user := &user.User{}
	row := db.QueryRow(ctx, SQL, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
