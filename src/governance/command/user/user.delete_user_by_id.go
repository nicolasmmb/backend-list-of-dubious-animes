package user

import "github.com/google/uuid"

type CommandDeleteUserById struct {
	ID uuid.UUID
}

func (c CommandDeleteUserById) IsCommand() {}

func (c CommandDeleteUserById) Data() any {
	return &c
}
