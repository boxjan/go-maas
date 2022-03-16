package maas

type Interface struct {
	Obj

	SystemId        string        `json:"system_id"`
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	Vlan            Vlan          `json:"vlan"`
	Type            string        `json:"type"`
	MacAddress      string        `json:"mac_address"`
	Parents         []string      `json:"parents"`
	Children        []string      `json:"children"`
	Tags            []string      `json:"tags"`
	Enabled         bool          `json:"enabled"`
	Links           UndefinedType `json:"links"`
	Params          UndefinedType `json:"params"`
	Discovered      UndefinedType `json:"discovered"`
	EffectiveMtu    int           `json:"effective_mtu"`
	Vendor          UndefinedType `json:"vendor"`
	Product         string        `json:"product"`
	FirmwareVersion UndefinedType `json:"firmware_version"`
	LinkConnected   bool          `json:"link_connected"`
	InterfaceSpeed  int           `json:"interface_speed"`
	LinkSpeed       int           `json:"link_speed"`
	SriovMaxVf      UndefinedType `json:"sriov_max_vf"`
}
