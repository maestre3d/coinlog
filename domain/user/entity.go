package user

import (
	"time"

	"github.com/maestre3d/coinlog/customtype"
)

// User individual interacting the system.
type User struct {
	ID          string
	DisplayName string // req
	customtype.Auditable
}

func newUser(args CreateCommand) User {
	return User{
		ID:          args.ID,
		DisplayName: args.DisplayName,
		Auditable:   customtype.NewAuditable(),
	}
}

func (u *User) update(args UpdateCommand) {
	u.DisplayName = args.DisplayName
	u.Version += 1
	u.UpdatedAt = time.Now().UTC()
}
