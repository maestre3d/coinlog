package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditableMixin{},
	}
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			NotEmpty().
			Immutable(),
		field.String("display_name").
			NotEmpty(),
		field.String("linked_to_user").
			Optional(),
		field.String("image_url").
			Optional(),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("contacts").
			Unique().
			Required(),
		edge.From("linked_to", User.Type).
			Ref("contact_links").
			Unique().
			Field("linked_to_user").
			Annotations(
				entsql.Annotation{
					Table:        "",
					Charset:      "",
					Collation:    "",
					Default:      "",
					DefaultExpr:  "",
					DefaultExprs: nil,
					Options:      "",
					Size:         0,
					WithComments: nil,
					Incremental:  nil,
					OnDelete:     entsql.SetNull,
					Check:        "",
					Checks:       nil,
				}),
	}
}
