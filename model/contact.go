package model

import (
	"github.com/maestre3d/coinlog/domainutil"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/valueobject"
)

var NewContactFromSQL domainutil.ConvertFunc[*ent.Contact, entity.Contact] = func(src *ent.Contact) entity.Contact {
	if src == nil {
		return entity.Contact{}
	}
	return entity.Contact{
		ID:          src.ID,
		User:        NewUserFromSQL(src.Edges.Owner),
		LinkedTo:    NewUserFromSQL(src.Edges.LinkedTo),
		DisplayName: src.DisplayName,
		ImageURL:    src.ImageURL,
		Auditable: valueobject.Auditable{
			IsActive:  src.IsActive,
			Version:   src.Version,
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
	}
}
