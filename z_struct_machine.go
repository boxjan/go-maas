package maas

type Machine struct {
	Node
	Owner          string          `json:"owner"`
	OwnerData      UndefinedStruct `json:"owner_data"`
	Locked         bool            `json:"locked"`
	CacheSets      UndefinedType   `json:"cache_sets"`
	BCaches        UndefinedType   `json:"b_caches"`
	BiosBootMethod string          `json:"bios_boot_method"`
	BootInterface  *Interface      `json:"boot_interface"`
	MinHweKernel   string          `json:"min_hwe_kernel"`
	HweKernel      string          `json:"hwe_kernel"`
}

func (m *Machine) recursiveClient() {
	m.BootInterface.setClient(m.getClient())
}
