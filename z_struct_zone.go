package maas

type Zone struct {
	Obj

	Name        string `json:"name"`
	Description string `json:"description"`
	Id          int    `json:"id"`
}
