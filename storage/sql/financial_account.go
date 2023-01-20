package sql

import (
	"context"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/ent"
	entschema "github.com/maestre3d/coinlog/ent/financialaccount"
	"github.com/maestre3d/coinlog/ent/predicate"
	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/storage"
)

type FinancialAccountStorage struct {
	client *ent.Client
}

var _ financialaccount.Repository = FinancialAccountStorage{}

func NewFinancialAccountStorage(e *ent.Client) FinancialAccountStorage {
	return FinancialAccountStorage{
		client: e,
	}
}

func (f FinancialAccountStorage) Save(ctx context.Context, v financialaccount.FinancialAccount) error {
	return f.client.FinancialAccount.Create().
		SetID(v.ID).
		SetDisplayName(v.DisplayName).
		SetOwnerID(v.User.ID).
		SetBankName(v.BankName).
		SetAccountType(v.AccountType.String()).
		SetBalance(v.Balance).
		SetCurrencyCode(v.CurrencyCode).
		SetIsActive(v.IsActive).
		SetCreatedAt(v.CreatedAt).
		SetUpdatedAt(v.UpdatedAt).
		SetVersion(v.Version).
		OnConflictColumns(entschema.FieldID).
		UpdateNewValues().
		Exec(ctx)
}

var _ parser.ParseFunc[*ent.FinancialAccount, financialaccount.FinancialAccount] = newFinancialAccountFromEnt

func newFinancialAccountFromEnt(src *ent.FinancialAccount) financialaccount.FinancialAccount {
	if src == nil {
		return financialaccount.FinancialAccount{}
	}

	return financialaccount.FinancialAccount{
		ID:           src.ID,
		User:         newUserFromEnt(src.Edges.Owner),
		DisplayName:  src.DisplayName,
		BankName:     src.BankName,
		AccountType:  financialaccount.NewAccountType(src.AccountType),
		Balance:      src.Balance,
		CurrencyCode: src.CurrencyCode,
		Auditable: customtype.Auditable{
			IsActive:  src.IsActive,
			Version:   src.Version,
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
	}
}

func (f FinancialAccountStorage) Remove(ctx context.Context, id string) error {
	if err := f.client.FinancialAccount.DeleteOneID(id).Exec(ctx); err != nil && !ent.IsNotFound(err) {
		return err
	}
	return nil
}

func (f FinancialAccountStorage) Get(ctx context.Context, id string) (*financialaccount.FinancialAccount, error) {
	schema, err := f.client.FinancialAccount.Query().Where(entschema.ID(id)).WithOwner().Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	out := newFinancialAccountFromEnt(schema)
	return &out, nil
}

func (f FinancialAccountStorage) buildQueryFunc(pred ...predicate.FinancialAccount) querySQLFunc[*ent.FinancialAccount] {
	return func(ctx context.Context, limit, offset int) ([]*ent.FinancialAccount, error) {
		return f.client.FinancialAccount.Query().
			Where(pred...).
			Limit(limit).
			Offset(offset).
			WithOwner().
			All(ctx)
	}
}

func (f FinancialAccountStorage) find(ctx context.Context, cr storage.Criteria, pred ...predicate.FinancialAccount) (
	[]financialaccount.FinancialAccount, storage.PageToken, error) {
	return paginateSQLFunc(ctx, cr, newFinancialAccountFromEnt, f.buildQueryFunc(pred...))
}

func (f FinancialAccountStorage) Find(ctx context.Context, cr storage.Criteria) (
	items []financialaccount.FinancialAccount, nextPage storage.PageToken, err error) {
	return f.find(ctx, cr, entschema.IsActive(true))
}
