package repository

import (
	"context"

	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/valueobject"
)

type Contact interface {
	Repository[string, entity.Contact]
	GetUserContacts(ctx context.Context, criteria valueobject.Criteria, userID string) ([]entity.Contact, valueobject.PageToken, error)
}
