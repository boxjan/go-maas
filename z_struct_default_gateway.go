package maas

type Gateway struct {
	GatewayIP string `json:"gateway_ip"`
	LinkId    int64  `json:"link_id"`
}

type DefaultGateway struct {
	Ipv4 Gateway `json:"ipv4"`
	Ipv6 Gateway `json:"ipv6"`
}
