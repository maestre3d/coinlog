package storage

// Repository writes and reads data using an underlying storage engine.
type Repository[T any] interface {
	Writer[T]
	Reader[T]
}
