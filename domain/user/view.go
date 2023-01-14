package user

import (
	"github.com/maestre3d/coinlog/domain"
)

type View struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	domain.Auditable
}

func NewView(usr User) View {
	return View{
		ID:          usr.ID,
		DisplayName: usr.DisplayName,
		Auditable: domain.Auditable{
			IsActive:  usr.Auditable.IsActive,
			Version:   usr.Auditable.Version,
			CreatedAt: usr.Auditable.CreatedAt,
			UpdatedAt: usr.Auditable.UpdatedAt,
		},
	}
}
