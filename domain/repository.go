package domain

import "context"

type RepositoryWriter[T any] interface {
	Save(ctx context.Context, v T) error
	Remove(ctx context.Context, id string) error
}

type RepositoryReader[T any] interface {
	Get(ctx context.Context, id string) (*T, error)
	Find(ctx context.Context, cr Criteria) (items []T, nextPage PageToken, err error)
}

type Repository[T any] interface {
	RepositoryWriter[T]
	RepositoryReader[T]
}
