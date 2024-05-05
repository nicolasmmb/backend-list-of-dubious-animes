package user

import (
	entity "backend/src/governance/entity/user"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *RepositoryUser) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	db := r.GetDB()

	SQL := `SELECT id, name, email, password, avatar, created_at, updated_at, deleted_at FROM users WHERE id = $1`

	user := &entity.User{}
	row := db.QueryRow(ctx, SQL, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	switch err {
	case nil:
		return user, nil
	case pgx.ErrNoRows:
		return nil, errors.New("--> User not found")
	case context.Canceled:
		return nil, err
	default:
		return nil, err
	}
}
