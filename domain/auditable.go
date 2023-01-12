package domain

import "time"

// Auditable metadata fields used to keep track of interactions between the system and its data.
type Auditable struct {
	IsActive  bool      `json:"is_active,omitempty"`
	Version   uint32    `json:"version,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
