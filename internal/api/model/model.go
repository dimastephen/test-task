package model

type FormWithTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      int    `json:"expires,omitempty"`
	Type         string `json:"tokenType,omitempty"`
}
