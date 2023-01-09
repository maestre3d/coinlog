// Package view domain-anemic structures used to present system data to an external agent (i.e. user, another system)
// using language primitive types only.
//
// Also known as Data Transfer Object (DTO).
package view

import (
	"time"

	"github.com/maestre3d/coinlog/valueobject"
)

type metadata struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Version   uint32 `json:"version"`
}

func newMetadata(a valueobject.Auditable) metadata {
	return metadata{
		CreatedAt: a.CreatedAt.Format(time.RFC3339),
		UpdatedAt: a.UpdatedAt.Format(time.RFC3339),
		Version:   a.Version,
	}
}
