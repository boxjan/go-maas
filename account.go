package maas

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

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

func (c *Client) GetAllAuthorisationToken() ([]AuthorisationToken, error) {
	rsp, err := c.TurnResponse(c.Get("account", "list_authorisation_tokens", nil))
	if err != nil {
		return nil, err
	}

	res := make([]AuthorisationToken, 0, 2)
	if err := json.Unmarshal(rsp, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) CreateAuthorisation(at *AuthorisationToken) (*AuthorisationToken, error) {
	rsp, err := c.TurnResponse(
		c.Post("account", "create_authorisation_token", url.Values{"name": []string{at.Name}}, nil))
	if err != nil {
		return nil, fmt.Errorf("request create authorisation failed with err: %+v", err)
	}

	err = json.Unmarshal(rsp, at)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (c *Client) UpdateAuthorisationName(at *AuthorisationToken, name string) error {
	_, err := c.TurnResponse(
		c.Post("account", "update_token_name",
			url.Values{
				"token": []string{at.Join().Token},
				"name":  []string{name},
			}, nil))
	if err != nil {
		return fmt.Errorf("update token name failed with err: %+v", err)
	}
	return nil
}

func (c *Client) DeleteAuthorisation(at *AuthorisationToken) error {
	_, err := c.TurnResponse(
		c.Post("account", "delete_authorisation_token",
			url.Values{
				"token_key": []string{at.Split().TokenKey},
			}, nil))
	if err != nil {
		return fmt.Errorf("update token name failed with err: %+v", err)
	}
	return nil
}
