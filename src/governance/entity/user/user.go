package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserID uuid.UUID

// remove Password field from json
type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string `json:"-"`
	Avatar    *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
type UserOptionUpdate func(c *User) error

func WithName(name string) UserOptionUpdate {
	return func(c *User) error {
		c.Name = name
		return nil
	}
}

func WithEmail(email string) UserOptionUpdate {
	return func(c *User) error {
		c.Email = email
		return nil
	}
}

func WithPassword(password string) UserOptionUpdate {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return func(c *User) error {
		c.Password = string(hashedPassword)
		return err
	}
}

func NewInstance(name, email, password string, avatar *string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		Avatar:    avatar,
		DeletedAt: nil,
	}, nil
}

func (u *User) Update(opts ...UserOptionUpdate) error {
	for _, opt := range opts {
		err := opt(u)
		return err
	}
	return nil
}

func (u *User) HasValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}
