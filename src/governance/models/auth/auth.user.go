package auth

import (
	"backend/src/governance/error/auth"
)

type AuthUserModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u AuthUserModel) Validate() error {
	if u.Email == "" {
		return auth.ErrEmailIsRequired
	}
	if u.Password == "" {
		return auth.ErrPasswordIsRequired
	}
	return nil
}
