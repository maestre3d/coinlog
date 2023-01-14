package contact

import (
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/user"
)

type View struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"display_name"`
	User        user.View  `json:"user,omitempty"`
	LinkedTo    *user.View `json:"linked_to"`
	ImageURL    *string    `json:"image_url"`
	domain.Auditable
}

func NewView(ctc Contact) View {
	v := View{
		ID:          ctc.ID,
		DisplayName: ctc.DisplayName,
		User:        user.NewView(ctc.User),
		LinkedTo:    nil,
		ImageURL:    domain.PtrIfNotEmpty(ctc.ImageURL),
		Auditable: domain.Auditable{
			IsActive:  ctc.Auditable.IsActive,
			Version:   ctc.Auditable.Version,
			CreatedAt: ctc.Auditable.CreatedAt,
			UpdatedAt: ctc.Auditable.UpdatedAt,
		},
	}
	if ctc.LinkedTo != nil {
		v.LinkedTo = domain.PtrTo(user.NewView(*ctc.LinkedTo))
	}
	return v
}
