package maas

type NumaNode struct {
	Id          int        `json:"id"`
	Index       int        `json:"index"`
	Memory      int        `json:"memory"`
	Cores       []int      `json:"cores"`
	HugePageSet []HugePage `json:"hugepages_set"`
}
