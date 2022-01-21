package maas

type NodeType int

const (
	NodeTypeMachine NodeType = iota
	NodeTypeDevice
	NodeTypeRackController
	NodeTypeRegionController
	NodeTypeRegionAndRackController
)

var NodeTypeStrMap = map[NodeType]string{
	NodeTypeMachine:                 "Machine",
	NodeTypeDevice:                  "Device",
	NodeTypeRackController:          "Rack controller",
	NodeTypeRegionController:        "Region controller",
	NodeTypeRegionAndRackController: "Region and rack controller",
}

func (nt NodeType) String() string {
	if str, exist := NodeTypeStrMap[nt]; exist {
		return str
	}
	return ""
}

func NodeTypeStrCover(s string) NodeType {
	for k, v := range NodeTypeStrMap {
		if v == s {
			return k
		}
	}
	return -1
}
