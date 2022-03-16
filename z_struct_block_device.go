package maas

type BlockDevice struct {
	Obj

	SystemId           string      `json:"system_id"`
	Id                 int         `json:"id"`
	Name               string      `json:"name"`
	Uuid               string      `json:"uuid"`
	Type               string      `json:"type"`
	Path               string      `json:"path"`
	Model              string      `json:"model"`
	Serial             string      `json:"serial"`
	IdPath             string      `json:"id_path"`
	Size               int64       `json:"size"`
	BlockSize          int         `json:"block_size"`
	AvailableSize      int64       `json:"available_size"`
	UsedSize           int         `json:"used_size"`
	UsedFor            string      `json:"used_for"`
	Tags               []string    `json:"tags"`
	Filesystem         FileSystem  `json:"filesystem"`
	PartitionTableType string      `json:"partition_table_type"`
	Partitions         []Partition `json:"partitions"`
	FirmwareVersion    string      `json:"firmware_version"`
	StoragePool        string      `json:"storage_pool"`
	NumaNode           int         `json:"numa_node"`
}
