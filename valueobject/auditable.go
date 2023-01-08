package valueobject

import "time"

// Auditable metadata fields used to keep track of interactions between the system and its data.
type Auditable struct {
	IsActive  bool
	Version   uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
