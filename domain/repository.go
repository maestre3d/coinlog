package domain

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, v T) error
	Get(ctx context.Context, id string) (*T, error)
	Find(ctx context.Context, cr Criteria) (items []T, nextPage PageToken, err error)
	Remove(ctx context.Context, id string) error
}
