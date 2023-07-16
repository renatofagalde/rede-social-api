package model

type Token struct {
	IsSuccessful bool   `json:"isSuccessful"`
	StatusCode   int    `json:"statusCode"`
	AccessToken  string `json:"accessToken"`
}
