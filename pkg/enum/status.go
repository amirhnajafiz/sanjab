package enum

type Status int

const (
	PendingStatus Status = iota + 1
	StartedStatus
	FailureStatus
)

func (s Status) ToString() string {
	switch s {
	case PendingStatus:
		return "Pending"
	case StartedStatus:
		return "Started"
	case FailureStatus:
		return "Failed"
	default:
		return "Unknown"
	}
}
