package contact

import (
	"time"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/user"
)

// Contact organization or individual a User interacts with. Might be a User of the system.
type Contact struct {
	ID          string
	User        user.User  // FK ->  users, req (created by)
	LinkedTo    *user.User // FK ->  users (nullable)
	DisplayName string     // req
	ImageURL    string
	customtype.Auditable
}

func newContact(args CreateCommand) Contact {
	c := Contact{
		ID: args.ID,
		User: user.User{
			ID: args.UserID,
		},
		LinkedTo:    nil,
		DisplayName: args.DisplayName,
		ImageURL:    args.ImageURL,
		Auditable:   customtype.NewAuditable(),
	}

	if args.LinkedToID != "" {
		c.LinkedTo = &user.User{
			ID: args.LinkedToID,
		}
	}

	return c
}

func (c *Contact) update(args UpdateCommand) {
	if args.LinkedToID == "" {
		c.LinkedTo = nil
	} else {
		c.LinkedTo = &user.User{ID: args.LinkedToID}
	}
	c.DisplayName = args.DisplayName
	c.ImageURL = args.ImageURL
	c.Auditable.Version += 1
	c.Auditable.UpdatedAt = time.Now().UTC()
}
