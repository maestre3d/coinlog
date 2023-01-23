package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinancialAccount holds the schema definition for the FinancialAccount entity.
type FinancialAccount struct {
	ent.Schema
}

func (FinancialAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditableMixin{},
	}
}

// Fields of the FinancialAccount.
func (FinancialAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			NotEmpty().
			Immutable(),
		field.String("display_name").
			NotEmpty(),
		field.String("bank_name").
			Optional(),
		field.String("account_type").
			NotEmpty(),
		field.Float("balance"),
		field.String("currency_code").
			NotEmpty(),
	}
}

// Edges of the FinancialAccount.
func (FinancialAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("financial_accounts").
			Unique().
			Required().
			Immutable(),
		edge.To("cards", Card.Type),
	}
}
