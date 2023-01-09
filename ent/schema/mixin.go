package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// -------------------------------------------------
// Mixin definition

// AuditableMixin implements the ent.Mixin for sharing
// auditable (CDC) fields with package schemas.
type AuditableMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (AuditableMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_active"),
		field.Uint32("version"),
		field.Time("created_at").
			Immutable(),
		field.Time("updated_at"),
	}
}
