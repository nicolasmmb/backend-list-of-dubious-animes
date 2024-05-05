package user

import (
	entity "backend/src/governance/entity/user"
	"context"

	"github.com/niko-labs/libs-go/helper/paginator"
)

func (r *RepositoryUser) GetUserWithFilter(ctx context.Context, pagination paginator.Pagination) ([]*entity.User, *int, error) {
	db := r.GetDB()

	SQL := `
		WITH user_counts AS (
			SELECT COUNT(*) AS total_count
			FROM users
			WHERE deleted_at IS NULL
		)
		SELECT 
			id, name, email, avatar, created_at, updated_at, deleted_at, 
			total_count
		FROM 
			users
		CROSS JOIN 
			user_counts
		WHERE 
			deleted_at IS NULL
		ORDER BY 
			name DESC
		LIMIT $1
		OFFSET $2;
	`

	rows, _ := db.Query(ctx, SQL, pagination.Limit, pagination.Offset())

	defer rows.Close()
	err := rows.Err()
	if err != nil {
		return nil, nil, err
	}

	users := []*entity.User{}
	var total *int

	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &total)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, user)
	}
	return users, total, nil
}
