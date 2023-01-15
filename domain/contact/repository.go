package contact

import (
	"context"

	"github.com/maestre3d/coinlog/storage"
)

type Repository interface {
	storage.Repository[Contact]
	GetUserContacts(ctx context.Context, criteria storage.Criteria, userID string) ([]Contact, storage.PageToken, error)
}
