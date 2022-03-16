package maas

const ResourcesMachines = "machines"

type Machine struct {
	Node
	Owner                  string          `json:"owner"`
	OwnerData              UndefinedStruct `json:"owner_data"`
	Locked                 bool            `json:"locked"`
	CacheSets              UndefinedType   `json:"cache_sets"`
	BCaches                UndefinedType   `json:"b_caches"`
	BiosBootMethod         string          `json:"bios_boot_method"`
	BootInterface          *Interface      `json:"boot_interface"`
	MinHweKernel           string          `json:"min_hwe_kernel"`
	HweKernel              string          `json:"hwe_kernel"`
	Storage                float64         `json:"storage"`
	Status                 NodeStatus      `json:"status"`
	Pool                   ResourcePool    `json:"pool"`
	DisableIpv4            bool            `json:"disable_ipv4"`
	ConstraintsByType      UndefinedStruct `json:"constraints_by_type"`
	BootDisk               BlockDevice     `json:"boot_disk"`
	BlockDeviceSet         []BlockDevice   `json:"blockdevice_set"`
	PhysicalBlockDeviceSet []BlockDevice   `json:"physicalblockdevice_set"`
	VirtualBlockDeviceSet  []BlockDevice   `json:"virtualblockdevice_set"`
	VolumeGroups           UndefinedType   `json:"volume_groups"`
	Raids                  UndefinedType   `json:"raids"`
	StatusMessage          UndefinedType   `json:"status_message"`
	StatusName             UndefinedType   `json:"status_name"`
	SpecialFileSystems     UndefinedType   `json:"special_filesystems"`
	Pod                    UndefinedType   `json:"pod"`
	DefaultGateway         DefaultGateway  `json:"default_gateway"`
	NumaNodeSet            []NumaNode      `json:"numanode_set"`
}
