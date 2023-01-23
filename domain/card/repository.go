package card

import (
	"context"

	"github.com/maestre3d/coinlog/storage"
)

type Repository interface {
	storage.Repository[Card]
	GetUserCards(ctx context.Context, cr storage.Criteria, userID string) ([]Card, storage.PageToken, error)
}
