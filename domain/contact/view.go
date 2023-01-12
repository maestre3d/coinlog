package contact

import (
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/user"
)

type CreateArgs struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
	UserID      string `json:"user_id" validate:"required"`
	LinkedToID  string `json:"linked_to_id"`
	ImageURL    string `json:"image_url" validate:"omitempty,url"`
}

type UpdateArgs struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
	LinkedToID  string `json:"linked_to"`
	ImageURL    string `json:"image_url" validate:"omitempty,url"`
}

type View struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"display_name"`
	User        user.View  `json:"user"`
	LinkedTo    *user.View `json:"linked_to"`
	ImageURL    *string    `json:"image_url"`
	domain.Auditable
}

var _ domain.Nullable[View] = View{}

func (v View) PtrIfNotEmpty() *View {
	if v.ID == "" {
		return nil
	}
	return &v
}

func NewView(ctc Contact) View {
	return View{
		ID:          ctc.ID,
		DisplayName: ctc.DisplayName,
		User:        user.NewView(ctc.User),
		LinkedTo:    domain.PtrIfNotEmpty(user.NewView(ctc.LinkedTo)),
		ImageURL:    domain.PtrIfNotEmpty(ctc.ImageURL),
		Auditable: domain.Auditable{
			IsActive:  ctc.Auditable.IsActive,
			Version:   ctc.Auditable.Version,
			CreatedAt: ctc.Auditable.CreatedAt,
			UpdatedAt: ctc.Auditable.UpdatedAt,
		},
	}
}
