package customtype

import "time"

// Auditable metadata fields used to keep track of interactions between the system and its data.
//
// Eases complexity from scenarios where Change-Data-Capture (CDC) is required.
type Auditable struct {
	IsActive  bool      `json:"is_active,omitempty"`
	Version   uint32    `json:"version,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewAuditable allocates an Auditable instance using default values.
func NewAuditable() Auditable {
	createTime := time.Now().UTC()
	return Auditable{
		IsActive:  true,
		Version:   1,
		CreatedAt: createTime,
		UpdatedAt: createTime,
	}
}
