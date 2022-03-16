package maas

type Ipaddress struct {
	Obj
	AllocType     IpaddressType `json:"alloc_type"`
	AllocTypeName string        `json:"alloc_type_name"`
	Created       string        `json:"created"`
	ResourceURI   string        `json:"resource_uri"`
	IP            string        `json:"ip"`
	Subnet        Subnet        `json:"subnet"`
	InterfaceSet  []Interface   `json:"interface_set"`
	Owner         User          `json:"owner"`
}
