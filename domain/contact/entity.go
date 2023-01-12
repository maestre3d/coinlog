package contact

import (
	"time"

	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/user"
)

// Contact organization or individual a User interacts with. Might be a User of the system.
type Contact struct {
	ID          string
	User        user.User // FK ->  users, req (created by)
	LinkedTo    user.User // FK ->  users (nullable)
	DisplayName string    // req
	ImageURL    string
	domain.Auditable
}

var _ domain.Nullable[Contact] = &Contact{}

func newContact(args CreateArgs) Contact {
	return Contact{
		ID: args.ID,
		User: user.User{
			ID: args.UserID,
		},
		LinkedTo: user.User{
			ID: args.LinkedToID,
		},
		DisplayName: args.DisplayName,
		ImageURL:    "",
		Auditable: domain.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
}

func (c *Contact) PtrIfNotEmpty() *Contact {
	if c.ID == "" {
		return nil
	}
	return c
}

func (c *Contact) Update(args UpdateArgs) {
	c.DisplayName = args.DisplayName
	c.ImageURL = args.ImageURL
	c.LinkedTo.ID = args.LinkedToID
	c.Auditable.Version += 1
	c.Auditable.UpdatedAt = time.Now().UTC()
}
