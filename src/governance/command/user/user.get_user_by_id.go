package user

import "github.com/google/uuid"

type CommandGetUserById struct {
	ID uuid.UUID
}

func (c CommandGetUserById) IsCommand() {}

func (c CommandGetUserById) Data() any {
	return &c
}
