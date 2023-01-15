package user

import (
	"github.com/maestre3d/coinlog/customtype"
)

type View struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	customtype.Auditable
}

func NewView(usr User) View {
	return View{
		ID:          usr.ID,
		DisplayName: usr.DisplayName,
		Auditable: customtype.Auditable{
			IsActive:  usr.Auditable.IsActive,
			Version:   usr.Auditable.Version,
			CreatedAt: usr.Auditable.CreatedAt,
			UpdatedAt: usr.Auditable.UpdatedAt,
		},
	}
}
