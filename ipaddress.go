package maas

func (c *Client) GetIpAddresses(filters ...string) (*[]Ipaddress, error) {
	params, err := FilterToUrlParams(filters...)
	if err != nil {
		return nil, err
	}

	var rsp []byte
	rsp, err = c.TurnResponse(c.Get("ipaddresses", "", params))
	if err != nil {
		return nil, err
	}

	ipAddresses := make([]Ipaddress, 0, 2)
	err = Unmarshal(rsp, &ipAddresses)
	if err != nil {
		return nil, err
	}

	return &ipAddresses, err
}
