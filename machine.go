package maas

import (
	"encoding/json"
	"errors"
	"net/url"
)

func (c *Client) GetMachines(filters ...string) ([]*Machine, error) {
	if len(filters)%2 != 0 {
		return nil, errors.New("errors machines filter")
	}

	params := url.Values{}
	for i := 0; i < len(filters); i += 2 {
		params[filters[i]] = []string{filters[i+1]}
	}
	rsp, err := c.TurnResponse(c.Get("machines", "", params))
	if err != nil {
		return nil, err
	}
	res := make([]*Machine, 0, 2)
	json.Unmarshal(rsp, &res)
	for _, m := range res {
		m.setClient(c)
		m.recursiveClient()
	}
	return res, err
}

func (m *Machine) GetPowerParameters() (*Power, error) {

	c := m.getClient()
	if c == nil {
		return nil, ErrEmptyClient
	}
	rsp, err := c.TurnResponse(c.Get("machines/"+m.SystemId, "power_parameters", nil))
	if err != nil {
		return nil, err
	}

	p := &Power{}
	json.Unmarshal(rsp, p)
	return p, nil
}
