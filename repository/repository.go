// Package repository interfaces abstracting internal persistence concrete implementations.
package repository

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, v T) error
	Get(ctx context.Context)
	Search(ctx context.Context) (items []T, nextPage string, err error)
	Remove(ctx context.Context, v T) error
}
