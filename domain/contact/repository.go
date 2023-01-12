package contact

import (
	"context"

	"github.com/maestre3d/coinlog/domain"
)

type Repository interface {
	domain.Repository[Contact]
	GetUserContacts(ctx context.Context, criteria domain.Criteria, userID string) ([]Contact, domain.PageToken, error)
}
