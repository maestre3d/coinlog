package entity

import "github.com/maestre3d/coinlog/valueobject"

// User individual interacting the system.
type User struct {
	ID          string
	DisplayName string // req
	Auditable   valueobject.Auditable
}
