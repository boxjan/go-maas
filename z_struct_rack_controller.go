package maas

type RackController struct {
	Node
	Owner      string        `json:"owner"`
	ServiceSet []ServiceSet  `json:"service_set"`
	Version    UndefinedType `json:"version"`
}
