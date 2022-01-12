package maas

type NodeStatus int

const ( // https://github.com/maas/maas/blob/31a95206af/src/maasserver/enum.py
	NodeStatusNew NodeStatus = iota
	NodeStatusCommissioning
	NodeStatusFailedCommissioning
	NodeStatusMissing
	NodeStatusReady
	NodeStatusReserved
	NodeStatusDeployed
	NodeStatusRetired
	NodeStatusBroken
	NodeStatusDeploying
	NodeStatusAllocated
	NodeStatusFailedDeployment
	NodeStatusReleasing
	NodeStatusDiskErasing
	NodeStatusFailDiskErasing
	NodeStatusRescueMode
	NodeStatusEnteringRescueMode
	NodeStatusFailedEnteringRescueMode
	NodeStatusExitingRescueMode
	NodeStatusFailedExitingRescueMode
	NodeStatusTesting
	NodeStatusFailedTesting
)

var NodeStatusMappingStr = map[NodeStatus]string{
	NodeStatusNew:                      "New",
	NodeStatusCommissioning:            "Commissioning",
	NodeStatusFailedCommissioning:      "Failed commissioning",
	NodeStatusMissing:                  "Missing",
	NodeStatusReady:                    "Ready",
	NodeStatusReserved:                 "Reserved",
	NodeStatusDeployed:                 "Deployed",
	NodeStatusRetired:                  "Retired",
	NodeStatusBroken:                   "Broken",
	NodeStatusDeploying:                "Deploying",
	NodeStatusAllocated:                "Allocated",
	NodeStatusFailedDeployment:         "Failed deployment",
	NodeStatusReleasing:                "Releasing",
	NodeStatusDiskErasing:              "Disk erasing",
	NodeStatusFailDiskErasing:          "Failed disk erasing",
	NodeStatusRescueMode:               "Rescue mode",
	NodeStatusEnteringRescueMode:       "Entering rescue mode",
	NodeStatusFailedEnteringRescueMode: "Failed to enter rescue mode",
	NodeStatusExitingRescueMode:        "Exiting rescue mode",
	NodeStatusFailedExitingRescueMode:  "Failed to exit rescue mode",
	NodeStatusTesting:                  "Testing",
	NodeStatusFailedTesting:            "Failed testing",
}

func (ns NodeStatus) String() string {
	if str, exist := NodeStatusMappingStr[ns]; exist {
		return str
	}
	return ""
}

func NodeStatusStrCover(s string) NodeStatus {
	for k, v := range NodeStatusMappingStr {
		if v == s {
			return k
		}
	}
	return -1
}
