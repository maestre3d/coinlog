package view

// ConvertModelToEntityFunc converts E to T.
type ConvertModelToEntityFunc[E, T any] func(src E) T

// NewCollection converts src (slice of type E) to a slice of T.
func NewCollection[E, T any](src []E, convertFunc ConvertModelToEntityFunc[E, T]) []T {
	buf := make([]T, 0, len(src))
	for _, item := range src {
		buf = append(buf, convertFunc(item))
	}
	return buf
}
