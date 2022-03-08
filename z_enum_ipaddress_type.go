package maas

type IpaddressType int

const (
	IpaddressTypeAuto IpaddressType = iota
	IpaddressTypeSticky
	IpaddressTypeUserReserved = iota + 4
	IpaddressTypeDhcp
	IpaddressTypeDiscovered
)

var IpaddressTypeStrMap = map[IpaddressType]string{
	IpaddressTypeAuto:         "Automatic",
	IpaddressTypeSticky:       "Static",
	IpaddressTypeUserReserved: "User reserved",
	IpaddressTypeDhcp:         "DHCP",
	IpaddressTypeDiscovered:   "Observed",
}

func (it IpaddressType) String() string {
	if str, exist := IpaddressTypeStrMap[it]; exist {
		return str
	}
	return ""
}

func IpaddressTypeStrCover(s string) IpaddressType {
	for k, v := range IpaddressTypeStrMap {
		if v == s {
			return k
		}
	}
	return -1
}
