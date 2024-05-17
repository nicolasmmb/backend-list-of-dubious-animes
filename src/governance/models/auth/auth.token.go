package auth

type TokenModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenIsValid struct {
	AccessToken string `json:"access_token"`
}

type TokenIsValidOutput struct {
	AccessToken bool `json:"access_token"`
}
