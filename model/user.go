package model

import (
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/entity"
)

func NewUserFromSQL(src *ent.User, dst *entity.User) {
	dst.ID = src.ID
	dst.DisplayName = src.DisplayName
	dst.Auditable.IsActive = src.IsActive
	dst.Auditable.Version = src.Version
	dst.Auditable.CreatedAt = src.CreatedAt
	dst.Auditable.UpdatedAt = src.UpdatedAt
}
