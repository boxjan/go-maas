package maas

func (c *Client) DeleteMachine(systemId string) error {
	_, err := c.TurnResponse(c.Delete(ResourcesMachines))
	return err
}

// GetMachine Reads a node with the given system_id.
func (c *Client) GetMachine(systemId string) (*Machine, error) {
	rsp, err := c.TurnResponse(c.Get(ResourcesMachines+systemId, "", nil))
	if err != nil {
		return nil, err
	}
	var m Machine
	return &m, Unmarshal(rsp, &m)
}

// GetMachineDetails will seturns system details -- for example, LLDP and lshw XML dumps.
// Note that this is returned as BSON and not JSON. This is for efficiency,
// but mainly because JSON can't do binary content without applying additional encoding
// like base-64.
func (c *Client) GetMachineDetails(systemId string) ([]byte, error) {
	return c.TurnResponse(c.Get(ResourcesMachines+systemId, "details", nil))
}

// GetMachineCurtinConfig return the rendered curtin configuration for the machine.
func (c *Client) GetMachineCurtinConfig(systemId string) ([]byte, error) {
	return c.TurnResponse(
		c.Get(ResourcesMachines+systemId, "get_curtin_config", nil))
}

// GetMachineToken return the maas token for the machine.
func (c *Client) GetMachineToken(systemId string) (*AuthorisationToken, error) {
	rsp, err := c.TurnResponse(c.Get(ResourcesMachines+systemId, "get_token", nil))
	if err != nil {
		return nil, err
	}
	var a AuthorisationToken
	return &a, Unmarshal(rsp, &a)
}

// GetMachinePowerParameters Gets power parameters for a given system_id, if any.
// For some types of power control this will include private information such as passwords and secret keys.
func (c *Client) GetMachinePowerParameters(systemId string) (*PowerParameters, error) {
	rsp, err := c.TurnResponse(c.Get(ResourcesMachines+systemId, "power_parameters", nil))
	if err != nil {
		return nil, err
	}
	var p PowerParameters
	return &p, Unmarshal(rsp, &p)
}

// GetMachinePowerState Gets the power state of a given node. MAAS sends a request to the node's power controller,
// which asks it about the node's state. The reply to this could be delayed by up to 30 seconds
// while waiting for the power controller to respond. Use this method sparingly as it ties up an appserver
// thread while waiting.
func (c *Client) GetMachinePowerState(systemId string) (string, error) {
	rsp, err := c.TurnResponse(c.Get(ResourcesMachines+systemId, "query_power_state", nil))
	if err != nil {
		return "?", err
	}
	type status struct {
		State string `json:"state"`
	}
	var s status
	return s.State, Unmarshal(rsp, s)
}
