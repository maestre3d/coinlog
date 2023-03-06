package user

import (
	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain"
)

type Event struct {
	UserID      string
	DisplayName string
	Action      string
	customtype.Auditable
}

var _ domain.Event = Event{}

func NewEvent(v User, op domain.EventAction) Event {
	return Event{
		UserID:      v.ID,
		DisplayName: v.DisplayName,
		Action:      op.String(),
		Auditable:   v.Auditable,
	}
}

func (e Event) GetKey() string {
	return e.UserID
}
