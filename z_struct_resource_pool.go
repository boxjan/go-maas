package maas

type ResourcePool struct {
	Obj
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          int    `json:"id"`
}
