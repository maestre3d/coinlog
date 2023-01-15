package pointer

// PtrIfNotEmpty retrieves a pointer of v if value is different from the zero-value of T.
func PtrIfNotEmpty[T comparable](v T) *T {
	var zeroVal T
	if v == zeroVal {
		return nil
	}
	return &v
}

// PtrTo retrieves a pointer of v.
func PtrTo[T any](v T) *T {
	return &v
}
