package maas

import "encoding/json"

type Power struct {
	Raw          []byte `json:"-"`
	PowerAddress string `json:"power_address"`
	PowerUser    string `json:"power_user"`
	PowerPass    string `json:"power_pass"`
}

type powerMid struct {
	PowerAddress string `json:"power_address"`
	PowerUser    string `json:"power_user"`
	PowerPass    string `json:"power_pass"`
}

func (p *Power) UnmarshalJSON(data []byte) error {
	mid := &powerMid{}
	err := json.Unmarshal(data, mid)
	if err != nil {
		return err
	}

	p.Raw = data
	p.PowerAddress = mid.PowerAddress
	p.PowerUser = mid.PowerUser
	p.PowerPass = mid.PowerPass

	return nil
}

func (p Power) MarshalJSON() ([]byte, error) {
	return p.Raw, nil
}
