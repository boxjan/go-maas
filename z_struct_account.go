package maas

import "strings"

type AuthorisationToken struct {
	Name        string `json:"name"`
	Token       string `json:"token"`
	TokenKey    string `json:"token_key,omitempty"`
	TokenSecret string `json:"token_secret,omitempty"`
	ConsumerKey string `json:"consumer_key,omitempty"`
}

func (at *AuthorisationToken) Split() *AuthorisationToken {
	elements := strings.Split(at.Token, ":")
	if len(elements) != 3 {
		return at
	}

	at.ConsumerKey = elements[0]
	at.TokenKey = elements[1]
	at.TokenSecret = elements[2]

	return at
}

func (at *AuthorisationToken) Join() *AuthorisationToken {
	if len(at.ConsumerKey) == 0 || len(at.TokenKey) == 0 || len(at.TokenSecret) == 0 {
		return at
	}

	at.Token = at.ConsumerKey + ":" + at.TokenKey + ":" + at.TokenSecret

	return at
}
