package view

import "github.com/maestre3d/coinlog/entity"

type User struct {
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	metadata
}

func NewUser(src entity.User) User {
	return User{
		UserID:      src.ID,
		DisplayName: src.DisplayName,
		metadata:    newMetadata(src.Auditable),
	}
}
