package user

import (
	"time"

	"github.com/google/uuid"
)

type BaseUserReturnModel struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Avatar    *string    `json:"avatar"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeleredAt *time.Time `json:"deleted_at"`
}

func ToBaseUserReturnModel(id uuid.UUID, name, email string, avatar *string, createdAt, updatedAt, deletedAt *time.Time) *BaseUserReturnModel {
	return &BaseUserReturnModel{
		ID:        id,
		Name:      name,
		Email:     email,
		Avatar:    avatar,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeleredAt: deletedAt,
	}
}
