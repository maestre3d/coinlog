package entity

import (
	"time"

	"github.com/maestre3d/coinlog/exception"
	"github.com/maestre3d/coinlog/valueobject"
)

var (
	ErrContactNotFound = exception.ResourceNotFound{
		Resource: "contact",
	}
)

// Contact organization or individual a User interacts with. Might be a User of the system.
type Contact struct {
	ID          string
	User        User   // FK ->  users, req (created by)
	LinkedTo    User   // FK ->  users (nullable)
	DisplayName string // req
	ImageURL    string
	Auditable   valueobject.Auditable
}

type ContactArgs struct {
	ID          string `validate:"required"`
	DisplayName string `validate:"required,lte=96"`
	UserID      string `validate:"required"`
	LinkedToID  string
	ImageURL    string `validate:"omitempty,url"`
}

func NewContact(args ContactArgs) (Contact, error) {
	if err := validate.Struct(args); err != nil {
		return Contact{}, err
	}

	return Contact{
		ID: args.ID,
		User: User{
			ID: args.UserID,
		},
		LinkedTo: User{
			ID: args.LinkedToID,
		},
		DisplayName: args.DisplayName,
		Auditable: valueobject.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}, nil
}

func (u *Contact) Update(args ContactArgs) error {
	if err := validate.Struct(args); err != nil {
		return err
	}
	u.DisplayName = args.DisplayName
	u.ImageURL = args.ImageURL
	u.Auditable.Version++
	u.Auditable.UpdatedAt = time.Now().UTC()
	return nil
}
