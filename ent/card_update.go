// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maestre3d/coinlog/ent/card"
	"github.com/maestre3d/coinlog/ent/financialaccount"
	"github.com/maestre3d/coinlog/ent/predicate"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetIsActive sets the "is_active" field.
func (cu *CardUpdate) SetIsActive(b bool) *CardUpdate {
	cu.mutation.SetIsActive(b)
	return cu
}

// SetVersion sets the "version" field.
func (cu *CardUpdate) SetVersion(u uint32) *CardUpdate {
	cu.mutation.ResetVersion()
	cu.mutation.SetVersion(u)
	return cu
}

// AddVersion adds u to the "version" field.
func (cu *CardUpdate) AddVersion(u int32) *CardUpdate {
	cu.mutation.AddVersion(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CardUpdate) SetUpdatedAt(t time.Time) *CardUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetFinancialAccountID sets the "financial_account_id" field.
func (cu *CardUpdate) SetFinancialAccountID(s string) *CardUpdate {
	cu.mutation.SetFinancialAccountID(s)
	return cu
}

// SetNillableFinancialAccountID sets the "financial_account_id" field if the given value is not nil.
func (cu *CardUpdate) SetNillableFinancialAccountID(s *string) *CardUpdate {
	if s != nil {
		cu.SetFinancialAccountID(*s)
	}
	return cu
}

// ClearFinancialAccountID clears the value of the "financial_account_id" field.
func (cu *CardUpdate) ClearFinancialAccountID() *CardUpdate {
	cu.mutation.ClearFinancialAccountID()
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *CardUpdate) SetDisplayName(s string) *CardUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetBankName sets the "bank_name" field.
func (cu *CardUpdate) SetBankName(s string) *CardUpdate {
	cu.mutation.SetBankName(s)
	return cu
}

// SetNillableBankName sets the "bank_name" field if the given value is not nil.
func (cu *CardUpdate) SetNillableBankName(s *string) *CardUpdate {
	if s != nil {
		cu.SetBankName(*s)
	}
	return cu
}

// ClearBankName clears the value of the "bank_name" field.
func (cu *CardUpdate) ClearBankName() *CardUpdate {
	cu.mutation.ClearBankName()
	return cu
}

// SetLastDigits sets the "last_digits" field.
func (cu *CardUpdate) SetLastDigits(u uint16) *CardUpdate {
	cu.mutation.ResetLastDigits()
	cu.mutation.SetLastDigits(u)
	return cu
}

// SetNillableLastDigits sets the "last_digits" field if the given value is not nil.
func (cu *CardUpdate) SetNillableLastDigits(u *uint16) *CardUpdate {
	if u != nil {
		cu.SetLastDigits(*u)
	}
	return cu
}

// AddLastDigits adds u to the "last_digits" field.
func (cu *CardUpdate) AddLastDigits(u int16) *CardUpdate {
	cu.mutation.AddLastDigits(u)
	return cu
}

// ClearLastDigits clears the value of the "last_digits" field.
func (cu *CardUpdate) ClearLastDigits() *CardUpdate {
	cu.mutation.ClearLastDigits()
	return cu
}

// SetBalance sets the "balance" field.
func (cu *CardUpdate) SetBalance(f float64) *CardUpdate {
	cu.mutation.ResetBalance()
	cu.mutation.SetBalance(f)
	return cu
}

// AddBalance adds f to the "balance" field.
func (cu *CardUpdate) AddBalance(f float64) *CardUpdate {
	cu.mutation.AddBalance(f)
	return cu
}

// SetBalanceLimit sets the "balance_limit" field.
func (cu *CardUpdate) SetBalanceLimit(f float64) *CardUpdate {
	cu.mutation.ResetBalanceLimit()
	cu.mutation.SetBalanceLimit(f)
	return cu
}

// AddBalanceLimit adds f to the "balance_limit" field.
func (cu *CardUpdate) AddBalanceLimit(f float64) *CardUpdate {
	cu.mutation.AddBalanceLimit(f)
	return cu
}

// SetCurrencyCode sets the "currency_code" field.
func (cu *CardUpdate) SetCurrencyCode(s string) *CardUpdate {
	cu.mutation.SetCurrencyCode(s)
	return cu
}

// SetCardType sets the "card_type" field.
func (cu *CardUpdate) SetCardType(s string) *CardUpdate {
	cu.mutation.SetCardType(s)
	return cu
}

// SetFinancialAccount sets the "financial_account" edge to the FinancialAccount entity.
func (cu *CardUpdate) SetFinancialAccount(f *FinancialAccount) *CardUpdate {
	return cu.SetFinancialAccountID(f.ID)
}

// Mutation returns the CardMutation object of the builder.
func (cu *CardUpdate) Mutation() *CardMutation {
	return cu.mutation
}

// ClearFinancialAccount clears the "financial_account" edge to the FinancialAccount entity.
func (cu *CardUpdate) ClearFinancialAccount() *CardUpdate {
	cu.mutation.ClearFinancialAccount()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CardMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CardUpdate) check() error {
	if v, ok := cu.mutation.DisplayName(); ok {
		if err := card.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Card.display_name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.CurrencyCode(); ok {
		if err := card.CurrencyCodeValidator(v); err != nil {
			return &ValidationError{Name: "currency_code", err: fmt.Errorf(`ent: validator failed for field "Card.currency_code": %w`, err)}
		}
	}
	if v, ok := cu.mutation.CardType(); ok {
		if err := card.CardTypeValidator(v); err != nil {
			return &ValidationError{Name: "card_type", err: fmt.Errorf(`ent: validator failed for field "Card.card_type": %w`, err)}
		}
	}
	if _, ok := cu.mutation.UserID(); cu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Card.user"`)
	}
	return nil
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: card.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.IsActive(); ok {
		_spec.SetField(card.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Version(); ok {
		_spec.SetField(card.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedVersion(); ok {
		_spec.AddField(card.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(card.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.SetField(card.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cu.mutation.BankName(); ok {
		_spec.SetField(card.FieldBankName, field.TypeString, value)
	}
	if cu.mutation.BankNameCleared() {
		_spec.ClearField(card.FieldBankName, field.TypeString)
	}
	if value, ok := cu.mutation.LastDigits(); ok {
		_spec.SetField(card.FieldLastDigits, field.TypeUint16, value)
	}
	if value, ok := cu.mutation.AddedLastDigits(); ok {
		_spec.AddField(card.FieldLastDigits, field.TypeUint16, value)
	}
	if cu.mutation.LastDigitsCleared() {
		_spec.ClearField(card.FieldLastDigits, field.TypeUint16)
	}
	if value, ok := cu.mutation.Balance(); ok {
		_spec.SetField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedBalance(); ok {
		_spec.AddField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.BalanceLimit(); ok {
		_spec.SetField(card.FieldBalanceLimit, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedBalanceLimit(); ok {
		_spec.AddField(card.FieldBalanceLimit, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.CurrencyCode(); ok {
		_spec.SetField(card.FieldCurrencyCode, field.TypeString, value)
	}
	if value, ok := cu.mutation.CardType(); ok {
		_spec.SetField(card.FieldCardType, field.TypeString, value)
	}
	if cu.mutation.FinancialAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.FinancialAccountTable,
			Columns: []string{card.FinancialAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: financialaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FinancialAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.FinancialAccountTable,
			Columns: []string{card.FinancialAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: financialaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetIsActive sets the "is_active" field.
func (cuo *CardUpdateOne) SetIsActive(b bool) *CardUpdateOne {
	cuo.mutation.SetIsActive(b)
	return cuo
}

// SetVersion sets the "version" field.
func (cuo *CardUpdateOne) SetVersion(u uint32) *CardUpdateOne {
	cuo.mutation.ResetVersion()
	cuo.mutation.SetVersion(u)
	return cuo
}

// AddVersion adds u to the "version" field.
func (cuo *CardUpdateOne) AddVersion(u int32) *CardUpdateOne {
	cuo.mutation.AddVersion(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CardUpdateOne) SetUpdatedAt(t time.Time) *CardUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetFinancialAccountID sets the "financial_account_id" field.
func (cuo *CardUpdateOne) SetFinancialAccountID(s string) *CardUpdateOne {
	cuo.mutation.SetFinancialAccountID(s)
	return cuo
}

// SetNillableFinancialAccountID sets the "financial_account_id" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableFinancialAccountID(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetFinancialAccountID(*s)
	}
	return cuo
}

// ClearFinancialAccountID clears the value of the "financial_account_id" field.
func (cuo *CardUpdateOne) ClearFinancialAccountID() *CardUpdateOne {
	cuo.mutation.ClearFinancialAccountID()
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *CardUpdateOne) SetDisplayName(s string) *CardUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetBankName sets the "bank_name" field.
func (cuo *CardUpdateOne) SetBankName(s string) *CardUpdateOne {
	cuo.mutation.SetBankName(s)
	return cuo
}

// SetNillableBankName sets the "bank_name" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableBankName(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetBankName(*s)
	}
	return cuo
}

// ClearBankName clears the value of the "bank_name" field.
func (cuo *CardUpdateOne) ClearBankName() *CardUpdateOne {
	cuo.mutation.ClearBankName()
	return cuo
}

// SetLastDigits sets the "last_digits" field.
func (cuo *CardUpdateOne) SetLastDigits(u uint16) *CardUpdateOne {
	cuo.mutation.ResetLastDigits()
	cuo.mutation.SetLastDigits(u)
	return cuo
}

// SetNillableLastDigits sets the "last_digits" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableLastDigits(u *uint16) *CardUpdateOne {
	if u != nil {
		cuo.SetLastDigits(*u)
	}
	return cuo
}

// AddLastDigits adds u to the "last_digits" field.
func (cuo *CardUpdateOne) AddLastDigits(u int16) *CardUpdateOne {
	cuo.mutation.AddLastDigits(u)
	return cuo
}

// ClearLastDigits clears the value of the "last_digits" field.
func (cuo *CardUpdateOne) ClearLastDigits() *CardUpdateOne {
	cuo.mutation.ClearLastDigits()
	return cuo
}

// SetBalance sets the "balance" field.
func (cuo *CardUpdateOne) SetBalance(f float64) *CardUpdateOne {
	cuo.mutation.ResetBalance()
	cuo.mutation.SetBalance(f)
	return cuo
}

// AddBalance adds f to the "balance" field.
func (cuo *CardUpdateOne) AddBalance(f float64) *CardUpdateOne {
	cuo.mutation.AddBalance(f)
	return cuo
}

// SetBalanceLimit sets the "balance_limit" field.
func (cuo *CardUpdateOne) SetBalanceLimit(f float64) *CardUpdateOne {
	cuo.mutation.ResetBalanceLimit()
	cuo.mutation.SetBalanceLimit(f)
	return cuo
}

// AddBalanceLimit adds f to the "balance_limit" field.
func (cuo *CardUpdateOne) AddBalanceLimit(f float64) *CardUpdateOne {
	cuo.mutation.AddBalanceLimit(f)
	return cuo
}

// SetCurrencyCode sets the "currency_code" field.
func (cuo *CardUpdateOne) SetCurrencyCode(s string) *CardUpdateOne {
	cuo.mutation.SetCurrencyCode(s)
	return cuo
}

// SetCardType sets the "card_type" field.
func (cuo *CardUpdateOne) SetCardType(s string) *CardUpdateOne {
	cuo.mutation.SetCardType(s)
	return cuo
}

// SetFinancialAccount sets the "financial_account" edge to the FinancialAccount entity.
func (cuo *CardUpdateOne) SetFinancialAccount(f *FinancialAccount) *CardUpdateOne {
	return cuo.SetFinancialAccountID(f.ID)
}

// Mutation returns the CardMutation object of the builder.
func (cuo *CardUpdateOne) Mutation() *CardMutation {
	return cuo.mutation
}

// ClearFinancialAccount clears the "financial_account" edge to the FinancialAccount entity.
func (cuo *CardUpdateOne) ClearFinancialAccount() *CardUpdateOne {
	cuo.mutation.ClearFinancialAccount()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Card entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks[*Card, CardMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CardUpdateOne) check() error {
	if v, ok := cuo.mutation.DisplayName(); ok {
		if err := card.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Card.display_name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.CurrencyCode(); ok {
		if err := card.CurrencyCodeValidator(v); err != nil {
			return &ValidationError{Name: "currency_code", err: fmt.Errorf(`ent: validator failed for field "Card.currency_code": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.CardType(); ok {
		if err := card.CardTypeValidator(v); err != nil {
			return &ValidationError{Name: "card_type", err: fmt.Errorf(`ent: validator failed for field "Card.card_type": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.UserID(); cuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Card.user"`)
	}
	return nil
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: card.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.IsActive(); ok {
		_spec.SetField(card.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Version(); ok {
		_spec.SetField(card.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedVersion(); ok {
		_spec.AddField(card.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(card.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.SetField(card.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.BankName(); ok {
		_spec.SetField(card.FieldBankName, field.TypeString, value)
	}
	if cuo.mutation.BankNameCleared() {
		_spec.ClearField(card.FieldBankName, field.TypeString)
	}
	if value, ok := cuo.mutation.LastDigits(); ok {
		_spec.SetField(card.FieldLastDigits, field.TypeUint16, value)
	}
	if value, ok := cuo.mutation.AddedLastDigits(); ok {
		_spec.AddField(card.FieldLastDigits, field.TypeUint16, value)
	}
	if cuo.mutation.LastDigitsCleared() {
		_spec.ClearField(card.FieldLastDigits, field.TypeUint16)
	}
	if value, ok := cuo.mutation.Balance(); ok {
		_spec.SetField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedBalance(); ok {
		_spec.AddField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.BalanceLimit(); ok {
		_spec.SetField(card.FieldBalanceLimit, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedBalanceLimit(); ok {
		_spec.AddField(card.FieldBalanceLimit, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.CurrencyCode(); ok {
		_spec.SetField(card.FieldCurrencyCode, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CardType(); ok {
		_spec.SetField(card.FieldCardType, field.TypeString, value)
	}
	if cuo.mutation.FinancialAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.FinancialAccountTable,
			Columns: []string{card.FinancialAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: financialaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FinancialAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.FinancialAccountTable,
			Columns: []string{card.FinancialAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: financialaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Card{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
