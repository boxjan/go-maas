package maas

type Node struct {
	Obj

	SystemId                     string       `json:"system_id"`
	Hostname                     string       `json:"hostname"`
	Description                  string       `json:"description"`
	HardwareUuid                 string       `json:"hardware_uuid"`
	Domain                       Domain       `json:"domain"`
	FQDN                         string       `json:"fqdn"`
	Architecture                 string       `json:"architecture"`
	CpuCount                     int64        `json:"cpu_count"`
	CpuSpeed                     int64        `json:"cpu_speed"`
	Memory                       int64        `json:"memory"`
	SwapSize                     int64        `json:"swap_size"`
	OSystem                      string       `json:"osystem"`
	DistroSeries                 string       `json:"distro_series"`
	NetBoot                      bool         `json:"netboot"`
	PowerType                    string       `json:"power_type"`
	PowerState                   string       `json:"power_state"`
	IpAddresses                  []string     `json:"ip_addresses"`
	InterfaceSet                 []Interface  `json:"interface_set"`
	Zone                         *Zone        `json:"zone"`
	StatusAction                 string       `json:"status_action"`
	NodeType                     NodeType     `json:"node_type"`
	NodeTypeName                 string       `json:"node_type_name"`
	CurrentCommissioningResultId int          `json:"current_commissioning_result_id"`
	CurrentTestingResultId       int          `json:"current_testing_result_id"`
	CurrentInstallationResultId  int          `json:"current_installation_result_id"`
	CommissioningStatus          ScriptStatus `json:"commissioning_status"`
	CommissioningStatusName      string       `json:"commissioning_status_name"`
	TestingStatus                int          `json:"testing_status"`
	TestingStatusName            string       `json:"testing_status_name"`
	CpuTestStatus                ScriptStatus `json:"cpu_test_status"`
	CpuTestStatusName            string       `json:"cpu_test_status_name"`
	MemoryTestStatus             ScriptStatus `json:"memory_test_status"`
	MemoryTestStatusName         string       `json:"memory_test_status_name"`
	NetworkTestStatus            ScriptStatus `json:"network_test_status"`
	NetworkTestStatusName        string       `json:"network_test_status_name"`
	StorageTestStatus            ScriptStatus `json:"storage_test_status"`
	StorageTestStatusName        string       `json:"storage_test_status_name"`
	OtherTestStatus              ScriptStatus `json:"other_test_status"`
	OtherTestStatusName          string       `json:"other_test_status_name"`
	HardwareInfo                 HardwareInfo `json:"hardware_info"`
	TagNames                     []string     `json:"tag_names"`
	InterfaceTestStatus          ScriptStatus `json:"interface_test_status"`
	InterfaceTestStatusName      string       `json:"interface_test_status_name"`
}
