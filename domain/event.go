package domain

import "fmt"

type Event interface {
	GetKey() string
}

type EventAction int

var _ fmt.Stringer = EventAction(0)

const (
	UnknownAction EventAction = iota
	CreateAction
	UpdateAction
	HardDeleteAction
	SoftDeleteAction
)

var EventActionMap = map[EventAction]string{
	UnknownAction:    "UNKNOWN",
	CreateAction:     "CREATE",
	UpdateAction:     "UPDATE",
	HardDeleteAction: "HARD_DELETE",
	SoftDeleteAction: "SOFT_DELETE",
}

var EventActionMapStr = map[string]EventAction{
	"UNKNOWN":     UnknownAction,
	"CREATE":      CreateAction,
	"UPDATE":      UpdateAction,
	"HARD_DELETE": HardDeleteAction,
	"SOFT_DELETE": SoftDeleteAction,
}

func (e EventAction) String() string {
	return EventActionMap[e]
}
