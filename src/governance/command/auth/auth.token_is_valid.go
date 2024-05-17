package auth

type CommandTokenIsValid struct {
	AccessToken  string

}

func (c CommandTokenIsValid) IsCommand() {}

func (c CommandTokenIsValid) Data() any {
	return &c
}
