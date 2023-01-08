package entity

import (
	"time"

	"github.com/maestre3d/coinlog/valueobject"
)

// User individual interacting the system.
type User struct {
	ID          string
	DisplayName string // req
	Auditable   valueobject.Auditable
}

type NewUserArgs struct {
	ID          string `validate:"required"`
	DisplayName string `validate:"required,lte=96"`
}

func NewUser(args NewUserArgs) (User, error) {
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
