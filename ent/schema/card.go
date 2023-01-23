package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

func (c Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditableMixin{},
	}
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			NotEmpty().
			Immutable(),
		field.String("financial_account_id").
			Optional(),
		field.String("display_name").
			NotEmpty(),
		field.String("bank_name").
			Optional(),
		field.Uint16("last_digits").
			Optional(),
		field.Float("balance"),
		field.Float("balance_limit"),
		field.String("currency_code").
			NotEmpty(),
		field.String("card_type").
			NotEmpty(),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("cards").
			Unique().
			Required().
			Immutable(),
		edge.From("financial_account", FinancialAccount.Type).
			Ref("cards").
			Unique().
			Field("financial_account_id"),
	}
}
