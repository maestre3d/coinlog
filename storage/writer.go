package storage

import "context"

// Writer writes data into an underlying storage engine.
type Writer[T any] interface {
	Save(ctx context.Context, v T) error
	Remove(ctx context.Context, id string) error
}
