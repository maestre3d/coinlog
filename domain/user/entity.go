package user

import (
	"time"

	"github.com/maestre3d/coinlog/domain"
)

// User individual interacting the system.
type User struct {
	ID          string
	DisplayName string // req
	domain.Auditable
}

var _ domain.Nullable[User] = &User{}

func newUser(args CreateArgs) User {
	return User{
		ID:          args.ID,
		DisplayName: args.DisplayName,
		Auditable: domain.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
}

func (u *User) PtrIfNotEmpty() *User {
	if u.ID == "" {
		return nil
	}
	return u
}

func (u *User) Update() {
	u.Version += 1
	u.UpdatedAt = time.Now().UTC()
}