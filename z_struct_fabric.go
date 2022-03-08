package maas

type Fabric struct {
	Obj

	Id        int    `json:"id"`
	Name      string `json:"name"`
	ClassType string `json:"class_type"`
	Vlans     []Vlan `json:"vlans"`
}

func (f *Fabric) recursiveClient() {
	for _, k := range f.Vlans {
		k.setClient(f.getClient())
	}
}
