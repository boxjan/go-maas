package maas

type PowerParameters struct {
	Raw          []byte `json:"-"`
	PowerAddress string `json:"power_address"`
	PowerUser    string `json:"power_user"`
	PowerPass    string `json:"power_pass"`
}

type powerParametersMid struct {
	PowerAddress string `json:"power_address"`
	PowerUser    string `json:"power_user"`
	PowerPass    string `json:"power_pass"`
}

func (p *PowerParameters) UnmarshalJSON(data []byte) error {
	mid := &powerParametersMid{}
	err := Unmarshal(data, mid)
	if err != nil {
		return err
	}

	p.Raw = data
	p.PowerAddress = mid.PowerAddress
	p.PowerUser = mid.PowerUser
	p.PowerPass = mid.PowerPass

	return nil
}

func (p PowerParameters) MarshalJSON() ([]byte, error) {
	if p.Raw != nil {
		return p.Raw, nil
	}
	mid := &powerParametersMid{}
	mid.PowerAddress = p.PowerAddress
	mid.PowerUser = p.PowerUser
	mid.PowerPass = p.PowerPass
	return Marshal(mid)
}
