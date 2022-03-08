package maas

type Fabric struct {
	Obj

	Id        int    `json:"id"`
	Name      string `json:"name"`
	ClassType string `json:"class_type"`
	Vlans     []Vlan `json:"vlans"`
}
