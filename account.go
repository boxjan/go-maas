package maas

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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
