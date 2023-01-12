package identifier

import "github.com/segmentio/ksuid"

var _ FactoryFunc = NewKSUID

// NewKSUID builds a K-Sortable Globally Unique ID (KSUID) with current time.
func NewKSUID() (string, error) {
	return ksuid.New().String(), nil
}
