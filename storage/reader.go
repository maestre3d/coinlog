package storage

import "context"

// Reader reads data from an underlying storage engine.
type Reader[T any] interface {
	Get(ctx context.Context, id string) (*T, error)
	Find(ctx context.Context, cr Criteria) (items []T, nextPage PageToken, err error)
}
