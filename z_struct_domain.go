package maas

type Domain struct {
	Obj

	Authoritative       bool        `json:"authoritative"`
	Ttl                 interface{} `json:"ttl"`
	Id                  int         `json:"id"`
	ResourceRecordCount int         `json:"resource_record_count"`
	IsDefault           bool        `json:"is_default"`
	Name                string      `json:"name"`
}
