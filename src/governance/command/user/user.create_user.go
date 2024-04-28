package user

type CommandCreateUser struct {
	Name     string
	Email    string
	Password string
	Avatar   *string
}

func (c CommandCreateUser) IsCommand() {}

func (c CommandCreateUser) Data() any {
	return &c
}
