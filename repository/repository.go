// Package repository interfaces abstracting internal persistence concrete implementations.
package repository

import (
	"context"
	"github.com/maestre3d/coinlog/valueobject"
)

type Repository[T any] interface {
	Save(ctx context.Context, v T) error
	Get(ctx context.Context, v *T) (found bool, err error)
	Search(ctx context.Context, c valueobject.Criteria) (items []T, nextPage valueobject.PageToken, err error)
	Remove(ctx context.Context, v T) error
}
