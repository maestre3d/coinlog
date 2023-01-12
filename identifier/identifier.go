package identifier

// FactoryFunc builds a unique identifier.
type FactoryFunc func() (string, error)
