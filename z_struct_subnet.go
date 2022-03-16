package maas

type Subnet struct {
	Obj
	Name            string   `json:"name,omitempty"`
	VLAN            Vlan     `json:"vlan,omitempty"`
	CIDR            string   `json:"cidr,omitempty"`
	RDNSMode        int      `json:"rdns_mode,omitempty"`
	GatewayIP       string   `json:"gateway_ip,omitempty"`
	DNSServers      []string `json:"dns_servers,omitempty"`
	AllowDNS        bool     `json:"allow_dns,omitempty"`
	AllowProxy      bool     `json:"allow_proxy,omitempty"`
	ActiveDiscovery bool     `json:"active_discovery,omitempty"`
	Managed         bool     `json:"managed,omitempty"`
	ID              int      `json:"id,omitempty"`
	Space           string   `json:"space,omitempty"`
}
