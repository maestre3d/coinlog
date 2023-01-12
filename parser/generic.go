package parser

// ParseFunc converts E to T.
type ParseFunc[E, T any] func(src E) T

// NewCollection converts src (slice of type E) to a slice of T.
func NewCollection[E, T any](src []E, convertFunc ParseFunc[E, T]) []T {
	if len(src) == 0 {
		return nil
	}

	buf := make([]T, 0, len(src))
	for _, item := range src {
		buf = append(buf, convertFunc(item))
	}
	return buf
}
