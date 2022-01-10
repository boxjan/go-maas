package maas

import "reflect"

type Node struct {
	Obj

	SystemId     string          `json:"system_id"`
	Hostname     string          `json:"hostname"`
	Description  string          `json:"description"`
	HardwareUuid string          `json:"hardware_uuid"`
	Domain       UndefinedStruct `json:"domain"`
	FQDN         string          `json:"fqdn"`
	Architecture string          `json:"architecture"`
	CpuCount     int64           `json:"cpu_count"`
	CpuSpeed     int64           `json:"cpu_speed"`
	Memory       int64           `json:"memory"`
	SwapSize     *int64          `json:"swap_size"`
}

func (n *Node) recursiveClient() {
	recursiveClient(n, reflect.TypeOf(&n))
}
