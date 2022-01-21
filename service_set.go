package maas

type ServiceSet struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	StatusInfo string `json:"status_info"`
}
