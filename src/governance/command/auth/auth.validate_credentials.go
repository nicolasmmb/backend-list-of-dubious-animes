package auth

type CommandAuthValidateCredentials struct {
	Email    string
	Password string
}

func (c CommandAuthValidateCredentials) IsCommand() {}

func (c CommandAuthValidateCredentials) Data() any {
	return &c
}
