package maas

type Vlan struct {
	Obj

	Id            int           `json:"id"`
	Name          string        `json:"name"`
	Vid           int           `json:"vid"`
	Fabric        string        `json:"fabric"`
	FabricId      int           `json:"fabric_id"`
	Mtu           int           `json:"mtu"`
	PrimaryRack   UndefinedType `json:"primary_rack"`
	SecondaryRack UndefinedType `json:"secondary_rack"`
	DhcpOn        bool          `json:"dhcp_on"`
	ExternalDhcp  UndefinedType `json:"external_dhcp"`
	RelayVlan     UndefinedType `json:"relay_vlan"`
	Space         string        `json:"space"`
}
