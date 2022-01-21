package maas

type BmcType int

const (
	BmcTypeBmc BmcType = iota
	BmcTypePod
)

var BmcTypeStrMap = map[BmcType]string{
	BmcTypeBmc: "BMC",
	BmcTypePod: "POD",
}

func (bt BmcType) String() string {
	if str, exist := BmcTypeStrMap[bt]; exist {
		return str
	}
	return ""
}

func BmcTypeStrCover(s string) BmcType {
	for k, v := range BmcTypeStrMap {
		if v == s {
			return k
		}
	}
	return -1
}
