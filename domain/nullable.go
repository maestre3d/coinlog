package domain

type Nullable[T any] interface {
	PtrIfNotEmpty() *T
}

func PtrIfNotEmpty[T comparable](v T) *T {
	var zeroVal T
	if v == zeroVal {
		return nil
	}
	return &v
}
