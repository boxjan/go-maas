package maas

type FileSystem struct {
	Fstype       string `json:"fstype"`
	Label        string `json:"label"`
	Uuid         string `json:"uuid"`
	MountPoint   string `json:"mount_point"`
	MountOptions string `json:"mount_options"`
}
