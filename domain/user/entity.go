package user

import (
	"time"

	"github.com/maestre3d/coinlog/domain"

	"github.com/maestre3d/coinlog/customtype"
)

// User individual interacting the system.
type User struct {
	ID          string
	DisplayName string // req
	customtype.Auditable
	*domain.Aggregate
}

func newUser(args CreateCommand) User {
	out := User{
		ID:          args.ID,
		DisplayName: args.DisplayName,
		Auditable:   customtype.NewAuditable(),
		Aggregate:   &domain.Aggregate{},
	}
	out.PushEvents(NewEvent(out, domain.CreateAction))
	return out
}

func (u *User) update(args UpdateCommand) {
	u.DisplayName = args.DisplayName
	u.Version += 1
	u.UpdatedAt = time.Now().UTC()
	u.PushEvents(NewEvent(*u, domain.UpdateAction))
}

func (u *User) delete(id string) {
	u.ID = id
	u.Aggregate = &domain.Aggregate{}
	u.PushEvents(NewEvent(*u, domain.HardDeleteAction))
}
