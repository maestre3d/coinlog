package contact

import (
	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/pointer"
)

type View struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"display_name"`
	User        user.View  `json:"user,omitempty"`
	LinkedTo    *user.View `json:"linked_to"`
	ImageURL    *string    `json:"image_url"`
	customtype.Auditable
}

func NewView(ctc Contact) View {
	v := View{
		ID:          ctc.ID,
		DisplayName: ctc.DisplayName,
		User:        user.NewView(ctc.User),
		LinkedTo:    nil,
		ImageURL:    pointer.PtrIfNotEmpty(ctc.ImageURL),
		Auditable: customtype.Auditable{
			IsActive:  ctc.Auditable.IsActive,
			Version:   ctc.Auditable.Version,
			CreatedAt: ctc.Auditable.CreatedAt,
			UpdatedAt: ctc.Auditable.UpdatedAt,
		},
	}
	if ctc.LinkedTo != nil {
		v.LinkedTo = pointer.PtrTo(user.NewView(*ctc.LinkedTo))
	}
	return v
}
