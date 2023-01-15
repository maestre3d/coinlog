package domain

import (
	"context"

	"github.com/maestre3d/coinlog/storage"
)

type BasicService[T any] interface {
	Create(ctx context.Context, args any) error
	Update(ctx context.Context, args any) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (T, error)
	List(ctx context.Context, cr storage.Criteria) ([]T, storage.PageToken, error)
}
