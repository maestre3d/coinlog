package domain

func PtrIfNotEmpty[T comparable](v T) *T {
	var zeroVal T
	if v == zeroVal {
		return nil
	}
	return &v
}

func PtrTo[T any](v T) *T {
	return &v
}
