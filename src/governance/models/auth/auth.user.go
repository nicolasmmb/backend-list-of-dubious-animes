package auth

import (
	"backend/src/governance/error/auth"
	"net/mail"
)

type AuthUserModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u AuthUserModel) Validate() error {
	if e, err := mail.ParseAddress(u.Email); err != nil || e.Address == "" {
		return auth.ErrEmailIsInvalid
	}
	if u.Password == "" {
		return auth.ErrPasswordIsRequired
	}
	return nil
}
