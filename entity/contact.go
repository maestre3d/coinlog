package entity

import "github.com/maestre3d/coinlog/valueobject"

// Contact organization or individual a User interacts with. Might be a User of the system.
type Contact struct {
	ID          string
	User        User   // FK ->  users, req (created by)
	LinkedTo    User   // FK ->  users (nullable)
	DisplayName string // req
	Auditable   valueobject.Auditable
}
