package maas

type ScriptStatus int

const (
	ScriptStatusPending ScriptStatus = iota
	ScriptStatusRunning
	ScriptStatusPassed
	ScriptStatusFailed
	ScriptStatusTimedout
	ScriptStatusAborted
	ScriptStatusDegraded
	ScriptStatusInstalling
	ScriptStatusFailedInstalling
	ScriptStatusSkipped
	ScriptStatusApplyingNetconf
	ScriptStatusFailedApplyingNetconf
)

var ScriptStatusStrMap = map[ScriptStatus]string{
	ScriptStatusPending:               "Pending",
	ScriptStatusRunning:               "Running",
	ScriptStatusPassed:                "Passed",
	ScriptStatusFailed:                "Failed",
	ScriptStatusTimedout:              "Timed out",
	ScriptStatusAborted:               "Aborted",
	ScriptStatusDegraded:              "Degraded",
	ScriptStatusInstalling:            "Installing dependencies",
	ScriptStatusFailedInstalling:      "Failed installing dependencies",
	ScriptStatusSkipped:               "Skipped",
	ScriptStatusApplyingNetconf:       "Applying custom network configuration",
	ScriptStatusFailedApplyingNetconf: "Failed to apply custom network configuration",
}

func (ss ScriptStatus) String() string {
	if str, exist := ScriptStatusStrMap[ss]; exist {
		return str
	}
	return ""
}

func ScriptStatusStrMappingStrCover(s string) ScriptStatus {
	for k, v := range ScriptStatusStrMap {
		if v == s {
			return k
		}
	}
	return -1
}
