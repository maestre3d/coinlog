package sql

import (
	"context"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/card"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/ent"
	entschema "github.com/maestre3d/coinlog/ent/card"
	"github.com/maestre3d/coinlog/ent/predicate"
	entuser "github.com/maestre3d/coinlog/ent/user"
	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/pointer"
	"github.com/maestre3d/coinlog/storage"
)

type CardStorage struct {
	db *ent.Client
}

var _ card.Repository = CardStorage{}

var _ parser.ParseFunc[*ent.Card, card.Card] = newCardFromEnt

func NewCardStorage(e *ent.Client) CardStorage {
	return CardStorage{
		db: e,
	}
}

func newCardFromEnt(src *ent.Card) card.Card {
	if src == nil {
		return card.Card{}
	}

	var acc *financialaccount.FinancialAccount
	if src.Edges.FinancialAccount != nil {
		acc = pointer.PtrTo(newFinancialAccountFromEnt(src.Edges.FinancialAccount))
	}
	return card.Card{
		ID:               src.ID,
		User:             newUserFromEnt(src.Edges.User),
		FinancialAccount: acc,
		DisplayName:      src.DisplayName,
		BankName:         src.BankName,
		LastDigits:       src.LastDigits,
		CardType:         card.NewTypeCard(src.CardType),
		Balance:          src.Balance,
		BalanceLimit:     src.BalanceLimit,
		CurrencyCode:     src.CurrencyCode,
		Auditable: customtype.Auditable{
			IsActive:  src.IsActive,
			Version:   src.Version,
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
	}
}

func (c CardStorage) create(ctx context.Context, v card.Card) error {
	stmt := c.db.Card.Create().
		SetID(v.ID).
		SetUserID(v.User.ID).
		SetDisplayName(v.DisplayName).
		SetBankName(v.BankName).
		SetLastDigits(v.LastDigits).
		SetCardType(v.CardType.String()).
		SetBalance(v.Balance).
		SetBalanceLimit(v.BalanceLimit).
		SetCurrencyCode(v.CurrencyCode).
		SetVersion(v.Version).
		SetCreatedAt(v.CreatedAt).
		SetUpdatedAt(v.UpdatedAt).
		SetIsActive(v.IsActive)

	if v.FinancialAccount != nil {
		stmt = stmt.SetFinancialAccountID(v.FinancialAccount.ID)
	}
	return stmt.Exec(ctx)
}

func (c CardStorage) update(ctx context.Context, v card.Card) error {
	stmt := c.db.Card.Update().
		Where(entschema.IDEQ(v.ID)).
		SetDisplayName(v.DisplayName).
		SetBankName(v.BankName).
		SetLastDigits(v.LastDigits).
		SetCardType(v.CardType.String()).
		SetBalance(v.Balance).
		SetBalanceLimit(v.BalanceLimit).
		SetCurrencyCode(v.CurrencyCode).
		SetVersion(v.Version).
		SetUpdatedAt(v.UpdatedAt).
		SetIsActive(v.IsActive)
	return stmt.Exec(ctx)
}

func (c CardStorage) Save(ctx context.Context, v card.Card) error {
	if v.Version == 1 {
		return c.create(ctx, v)
	}

	return c.update(ctx, v)
}

func (c CardStorage) Get(ctx context.Context, id string) (*card.Card, error) {
	out, err := c.db.Card.Query().Where(entschema.ID(id)).WithUser().WithFinancialAccount().Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	cp := newCardFromEnt(out)
	return &cp, nil
}

func (c CardStorage) buildQueryFunc(pred ...predicate.Card) querySQLFunc[*ent.Card] {
	return func(ctx context.Context, limit, offset int) ([]*ent.Card, error) {
		return c.db.Card.Query().
			Where(pred...).
			Limit(limit).
			Offset(offset).
			WithUser().
			WithFinancialAccount().
			All(ctx)
	}
}

func (c CardStorage) find(ctx context.Context, cr storage.Criteria, pred ...predicate.Card) ([]card.Card,
	storage.PageToken, error) {
	return paginateSQLFunc(ctx, cr, newCardFromEnt, c.buildQueryFunc(pred...))
}

func (c CardStorage) Find(ctx context.Context, cr storage.Criteria) (items []card.Card, nextPage storage.PageToken, err error) {
	return c.find(ctx, cr, entschema.IsActive(true))
}

func (c CardStorage) GetUserCards(ctx context.Context, cr storage.Criteria,
	userID string) ([]card.Card, storage.PageToken, error) {
	return c.find(ctx, cr, entschema.And(
		entschema.HasUserWith(entuser.ID(userID)),
		entschema.IsActive(true)),
	)
}

func (c CardStorage) Remove(ctx context.Context, id string) error {
	err := c.db.Card.DeleteOneID(id).Exec(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil
	}

	return err
}
