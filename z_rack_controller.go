package maas

type RackController struct {
	Node
	Owner string `json:"owner"`
}
