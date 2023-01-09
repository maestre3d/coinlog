package domainutil

// NewPtrTo allocates a new pointer to v if not empty.
func NewPtrTo[T comparable](v T) *T {
	var zeroVal T
	if v == zeroVal {
		return nil
	}
	return &v
}
