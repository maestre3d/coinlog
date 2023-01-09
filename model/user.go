package model

import (
	"github.com/maestre3d/coinlog/domainutil"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/valueobject"
)

var NewUserFromSQL domainutil.ConvertFunc[*ent.User, entity.User] = func(src *ent.User) entity.User {
	if src == nil {
		return entity.User{}
	}
	return entity.User{
		ID:          src.ID,
		DisplayName: src.DisplayName,
		Auditable: valueobject.Auditable{
			IsActive:  src.IsActive,
			Version:   src.Version,
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
	}
}
