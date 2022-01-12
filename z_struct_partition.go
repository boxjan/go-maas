package maas

type Partition struct {
	Obj

	SystemId   string     `json:"system_id"`
	DeviceId   int        `json:"device_id"`
	Id         int        `json:"id"`
	Uuid       string     `json:"uuid"`
	Path       string     `json:"path"`
	Type       string     `json:"type"`
	Size       int64      `json:"size"`
	BootAble   bool       `json:"bootable"`
	Tags       []string   `json:"tags"`
	UsedFor    string     `json:"used_for"`
	Filesystem FileSystem `json:"filesystem"`
}
