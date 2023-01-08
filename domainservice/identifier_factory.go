package domainservice

import "github.com/segmentio/ksuid"

// IdentifierFactoryFunc builds a unique identifier.
type IdentifierFactoryFunc func() (string, error)

// KSUIDFactory builds a K-Sortable Globally Unique ID (KSUID) with current time.
var KSUIDFactory IdentifierFactoryFunc = func() (string, error) {
	return ksuid.New().String(), nil
}
