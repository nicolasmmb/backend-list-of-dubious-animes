package user

import "backend/src/governance/error/user"

type UpdateUserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

func (u UpdateUserModel) Validate() error {
	if u.Name == "" {
		return user.ErrNameIsRequired
	}
	if u.Email == "" {
		return user.ErrEmailIsRequired
	}
	return nil
}

type UpdateUserDTOOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
