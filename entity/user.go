package entity

import (
	"github.com/maestre3d/coinlog/exception"
	"time"

	"github.com/maestre3d/coinlog/valueobject"
)

var (
	ErrUserNotFound = exception.ResourceNotFound{
		Resource: "user",
	}
)

// User individual interacting the system.
type User struct {
	ID          string
	DisplayName string // req
	Auditable   valueobject.Auditable
}

type UserArgs struct {
	ID          string `validate:"required"`
	DisplayName string `validate:"required,lte=96"`
}

func NewUser(args UserArgs) (User, error) {
	if err := validate.Struct(args); err != nil {
		return User{}, err
	}

	return User{
		ID:          args.ID,
		DisplayName: args.DisplayName,
		Auditable: valueobject.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}, nil
}

func (u *User) Update(args UserArgs) error {
	if err := validate.Struct(args); err != nil {
		return err
	}
	u.DisplayName = args.DisplayName
	u.Auditable.Version++
	u.Auditable.UpdatedAt = time.Now().UTC()
	return nil
}
