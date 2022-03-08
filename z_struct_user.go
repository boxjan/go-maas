package maas

type User struct {
	Obj

	IsSuperUser bool   `json:"is_superuser"`
	IsLocal     bool   `json:"is_local"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
}
