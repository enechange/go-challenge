package domain

type Status int

const (
	Available Status = iota + 1
	Blocked
	Charging
	Inoperative
	OutOfOrder
	Planned
	Removed
	Reserved
	Unknown
)

type EVSE struct {
	UID    string
	Status Status
}
