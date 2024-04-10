package constants

type Status int8

const (
    AVAILABLE Status = iota + 1
    BLOCKED
    CHARGING
    INOPERATIVE
    OUTOFORDER
    PLANNED
    REMOVED
    RESERVED
    UNKNOWN
)

func (s Status) String() string {
    return [...]string{"AVAILABLE", "BLOCKED", "CHARGING", "INOPERATIVE", "OUTOFORDER", "PLANNED", "REMOVED", "RESERVED", "UNKNOWN"}[s-1]
}