package user

import "backend/src/governance/error/user"

type CreateUserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u CreateUserModel) Validate() error {
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
