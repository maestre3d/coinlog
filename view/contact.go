package view

import (
	"github.com/maestre3d/coinlog/domainutil"
	"github.com/maestre3d/coinlog/entity"
)

type Contact struct {
	ContactID   string  `json:"contact_id"`
	DisplayName string  `json:"display_name"`
	Owner       User    `json:"owner"`
	LinkedTo    *User   `json:"linked_to"`
	ImageURL    *string `json:"image_url"`
	metadata
}

func NewContact(src entity.Contact) Contact {
	c := Contact{
		ContactID:   src.ID,
		DisplayName: src.DisplayName,
		Owner:       NewUser(src.User),
		LinkedTo:    nil,
		ImageURL:    domainutil.NewPtrTo(src.ImageURL),
		metadata:    newMetadata(src.Auditable),
	}
	if src.LinkedTo.ID != "" {
		linked := NewUser(src.LinkedTo)
		c.LinkedTo = &linked
	}
	return c
}
