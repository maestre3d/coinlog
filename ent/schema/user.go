package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditableMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			NotEmpty().
			Immutable(),
		field.String("display_name").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contacts", Contact.Type),
		edge.To("contact_links", Contact.Type),
	}
}
