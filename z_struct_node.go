package maas

type Node struct {
	Obj

	SystemId     string       `json:"system_id"`
	Hostname     string       `json:"hostname"`
	Description  string       `json:"description"`
	HardwareUuid string       `json:"hardware_uuid"`
	Domain       *Domain      `json:"domain,omitempty"`
	FQDN         string       `json:"fqdn"`
	Architecture string       `json:"architecture"`
	CpuCount     int64        `json:"cpu_count"`
	CpuSpeed     int64        `json:"cpu_speed"`
	Memory       int64        `json:"memory"`
	SwapSize     *int64       `json:"swap_size"`
	OSystem      string       `json:"osystem"`
	DistroSeries string       `json:"distro_series"`
	NetBoot      bool         `json:"netboot"`
	PowerType    string       `json:"power_type"`
	PowerState   string       `json:"power_state"`
	IpAddresses  []string     `json:"ip_addresses"`
	InterfaceSet []*Interface `json:"interface_set"`
	Zone         *Zone        `json:"zone"`
	StatusAction string       `json:"status_action"`
	NodeType     NodeType     `json:"node_type"`
	NodeTypeName string       `json:"node_type_name"`

	HardwareInfo HardwareInfo `json:"hardware_info"`
}

func (n *Node) recursiveClient() {
	for _, k := range n.InterfaceSet {
		k.setClient(n.getClient())
	}
	n.Domain.setClient(n.getClient())
}
