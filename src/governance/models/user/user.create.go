package user

import (
	"backend/src/governance/error/user"
	"net/mail"
)

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
	if e, err := mail.ParseAddress(u.Email); err != nil || e.Address == "" {
		return user.ErrEmailIsInvalid
	}
	if u.Password == "" {
		return user.ErrPasswordIsRequired
	}
	return nil
}
