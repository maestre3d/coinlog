package user

import (
	"github.com/maestre3d/coinlog/domain"
)

type CreateArgs struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
}

type UpdateArgs struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
}

type View struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	domain.Auditable
}

var _ domain.Nullable[View] = View{}

func (v View) PtrIfNotEmpty() *View {
	if v.ID == "" {
		return nil
	}
	return &v
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
