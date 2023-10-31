package enum

type Status int

const (
	PendingStatus Status = iota + 1
	StartedStatus
	FailureStatus
	DisableStatus
)

func (s Status) ToString() string {
	switch s {
	case PendingStatus:
		return "Pending"
	case StartedStatus:
		return "Started"
	case FailureStatus:
		return "Failed"
	case DisableStatus:
		return "Disable"
	default:
		return "Unknown"
	}
}
