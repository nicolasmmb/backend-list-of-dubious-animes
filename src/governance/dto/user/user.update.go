package user

import "backend/src/governance/error/user"

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u UpdateUserDTO) Validate() error {
	if u.Name == "" {
		return user.ErrNameIsRequired
	}
	if u.Email == "" {
		return user.ErrEmailIsRequired
	}
	if u.Password == "" {
		return user.ErrPasswordIsRequired
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
