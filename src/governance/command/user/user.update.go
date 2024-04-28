package user

type CommandUpdateUser struct {
	Name     string
	Email    string
	Password string
	Avatar   *string
}

func (c CommandUpdateUser) IsCommand() {}

func (c CommandUpdateUser) Data() any {
	return &c
}
