// Code generated by ent, DO NOT EDIT.

package card

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/maestre3d/coinlog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldID, id))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldIsActive, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v uint32) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldVersion, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldUpdatedAt, v))
}

// FinancialAccountID applies equality check predicate on the "financial_account_id" field. It's identical to FinancialAccountIDEQ.
func FinancialAccountID(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldFinancialAccountID, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldDisplayName, v))
}

// BankName applies equality check predicate on the "bank_name" field. It's identical to BankNameEQ.
func BankName(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBankName, v))
}

// LastDigits applies equality check predicate on the "last_digits" field. It's identical to LastDigitsEQ.
func LastDigits(v uint16) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldLastDigits, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBalance, v))
}

// BalanceLimit applies equality check predicate on the "balance_limit" field. It's identical to BalanceLimitEQ.
func BalanceLimit(v float64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBalanceLimit, v))
}

// CurrencyCode applies equality check predicate on the "currency_code" field. It's identical to CurrencyCodeEQ.
func CurrencyCode(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCurrencyCode, v))
}

// CardType applies equality check predicate on the "card_type" field. It's identical to CardTypeEQ.
func CardType(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCardType, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldIsActive, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v uint32) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v uint32) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...uint32) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...uint32) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v uint32) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v uint32) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v uint32) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v uint32) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldVersion, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldUpdatedAt, v))
}

// FinancialAccountIDEQ applies the EQ predicate on the "financial_account_id" field.
func FinancialAccountIDEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldFinancialAccountID, v))
}

// FinancialAccountIDNEQ applies the NEQ predicate on the "financial_account_id" field.
func FinancialAccountIDNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldFinancialAccountID, v))
}

// FinancialAccountIDIn applies the In predicate on the "financial_account_id" field.
func FinancialAccountIDIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldFinancialAccountID, vs...))
}

// FinancialAccountIDNotIn applies the NotIn predicate on the "financial_account_id" field.
func FinancialAccountIDNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldFinancialAccountID, vs...))
}

// FinancialAccountIDGT applies the GT predicate on the "financial_account_id" field.
func FinancialAccountIDGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldFinancialAccountID, v))
}

// FinancialAccountIDGTE applies the GTE predicate on the "financial_account_id" field.
func FinancialAccountIDGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldFinancialAccountID, v))
}

// FinancialAccountIDLT applies the LT predicate on the "financial_account_id" field.
func FinancialAccountIDLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldFinancialAccountID, v))
}

// FinancialAccountIDLTE applies the LTE predicate on the "financial_account_id" field.
func FinancialAccountIDLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldFinancialAccountID, v))
}

// FinancialAccountIDContains applies the Contains predicate on the "financial_account_id" field.
func FinancialAccountIDContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldFinancialAccountID, v))
}

// FinancialAccountIDHasPrefix applies the HasPrefix predicate on the "financial_account_id" field.
func FinancialAccountIDHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldFinancialAccountID, v))
}

// FinancialAccountIDHasSuffix applies the HasSuffix predicate on the "financial_account_id" field.
func FinancialAccountIDHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldFinancialAccountID, v))
}

// FinancialAccountIDIsNil applies the IsNil predicate on the "financial_account_id" field.
func FinancialAccountIDIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldFinancialAccountID))
}

// FinancialAccountIDNotNil applies the NotNil predicate on the "financial_account_id" field.
func FinancialAccountIDNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldFinancialAccountID))
}

// FinancialAccountIDEqualFold applies the EqualFold predicate on the "financial_account_id" field.
func FinancialAccountIDEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldFinancialAccountID, v))
}

// FinancialAccountIDContainsFold applies the ContainsFold predicate on the "financial_account_id" field.
func FinancialAccountIDContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldFinancialAccountID, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldDisplayName, v))
}

// BankNameEQ applies the EQ predicate on the "bank_name" field.
func BankNameEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBankName, v))
}

// BankNameNEQ applies the NEQ predicate on the "bank_name" field.
func BankNameNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldBankName, v))
}

// BankNameIn applies the In predicate on the "bank_name" field.
func BankNameIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldBankName, vs...))
}

// BankNameNotIn applies the NotIn predicate on the "bank_name" field.
func BankNameNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldBankName, vs...))
}

// BankNameGT applies the GT predicate on the "bank_name" field.
func BankNameGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldBankName, v))
}

// BankNameGTE applies the GTE predicate on the "bank_name" field.
func BankNameGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldBankName, v))
}

// BankNameLT applies the LT predicate on the "bank_name" field.
func BankNameLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldBankName, v))
}

// BankNameLTE applies the LTE predicate on the "bank_name" field.
func BankNameLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldBankName, v))
}

// BankNameContains applies the Contains predicate on the "bank_name" field.
func BankNameContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldBankName, v))
}

// BankNameHasPrefix applies the HasPrefix predicate on the "bank_name" field.
func BankNameHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldBankName, v))
}

// BankNameHasSuffix applies the HasSuffix predicate on the "bank_name" field.
func BankNameHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldBankName, v))
}

// BankNameIsNil applies the IsNil predicate on the "bank_name" field.
func BankNameIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldBankName))
}

// BankNameNotNil applies the NotNil predicate on the "bank_name" field.
func BankNameNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldBankName))
}

// BankNameEqualFold applies the EqualFold predicate on the "bank_name" field.
func BankNameEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldBankName, v))
}

// BankNameContainsFold applies the ContainsFold predicate on the "bank_name" field.
func BankNameContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldBankName, v))
}

// LastDigitsEQ applies the EQ predicate on the "last_digits" field.
func LastDigitsEQ(v uint16) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldLastDigits, v))
}

// LastDigitsNEQ applies the NEQ predicate on the "last_digits" field.
func LastDigitsNEQ(v uint16) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldLastDigits, v))
}

// LastDigitsIn applies the In predicate on the "last_digits" field.
func LastDigitsIn(vs ...uint16) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldLastDigits, vs...))
}

// LastDigitsNotIn applies the NotIn predicate on the "last_digits" field.
func LastDigitsNotIn(vs ...uint16) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldLastDigits, vs...))
}

// LastDigitsGT applies the GT predicate on the "last_digits" field.
func LastDigitsGT(v uint16) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldLastDigits, v))
}

// LastDigitsGTE applies the GTE predicate on the "last_digits" field.
func LastDigitsGTE(v uint16) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldLastDigits, v))
}

// LastDigitsLT applies the LT predicate on the "last_digits" field.
func LastDigitsLT(v uint16) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldLastDigits, v))
}

// LastDigitsLTE applies the LTE predicate on the "last_digits" field.
func LastDigitsLTE(v uint16) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldLastDigits, v))
}

// LastDigitsIsNil applies the IsNil predicate on the "last_digits" field.
func LastDigitsIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldLastDigits))
}

// LastDigitsNotNil applies the NotNil predicate on the "last_digits" field.
func LastDigitsNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldLastDigits))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldBalance, v))
}

// BalanceLimitEQ applies the EQ predicate on the "balance_limit" field.
func BalanceLimitEQ(v float64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBalanceLimit, v))
}

// BalanceLimitNEQ applies the NEQ predicate on the "balance_limit" field.
func BalanceLimitNEQ(v float64) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldBalanceLimit, v))
}

// BalanceLimitIn applies the In predicate on the "balance_limit" field.
func BalanceLimitIn(vs ...float64) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldBalanceLimit, vs...))
}

// BalanceLimitNotIn applies the NotIn predicate on the "balance_limit" field.
func BalanceLimitNotIn(vs ...float64) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldBalanceLimit, vs...))
}

// BalanceLimitGT applies the GT predicate on the "balance_limit" field.
func BalanceLimitGT(v float64) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldBalanceLimit, v))
}

// BalanceLimitGTE applies the GTE predicate on the "balance_limit" field.
func BalanceLimitGTE(v float64) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldBalanceLimit, v))
}

// BalanceLimitLT applies the LT predicate on the "balance_limit" field.
func BalanceLimitLT(v float64) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldBalanceLimit, v))
}

// BalanceLimitLTE applies the LTE predicate on the "balance_limit" field.
func BalanceLimitLTE(v float64) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldBalanceLimit, v))
}

// CurrencyCodeEQ applies the EQ predicate on the "currency_code" field.
func CurrencyCodeEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCurrencyCode, v))
}

// CurrencyCodeNEQ applies the NEQ predicate on the "currency_code" field.
func CurrencyCodeNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCurrencyCode, v))
}

// CurrencyCodeIn applies the In predicate on the "currency_code" field.
func CurrencyCodeIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCurrencyCode, vs...))
}

// CurrencyCodeNotIn applies the NotIn predicate on the "currency_code" field.
func CurrencyCodeNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCurrencyCode, vs...))
}

// CurrencyCodeGT applies the GT predicate on the "currency_code" field.
func CurrencyCodeGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCurrencyCode, v))
}

// CurrencyCodeGTE applies the GTE predicate on the "currency_code" field.
func CurrencyCodeGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCurrencyCode, v))
}

// CurrencyCodeLT applies the LT predicate on the "currency_code" field.
func CurrencyCodeLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCurrencyCode, v))
}

// CurrencyCodeLTE applies the LTE predicate on the "currency_code" field.
func CurrencyCodeLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCurrencyCode, v))
}

// CurrencyCodeContains applies the Contains predicate on the "currency_code" field.
func CurrencyCodeContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldCurrencyCode, v))
}

// CurrencyCodeHasPrefix applies the HasPrefix predicate on the "currency_code" field.
func CurrencyCodeHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldCurrencyCode, v))
}

// CurrencyCodeHasSuffix applies the HasSuffix predicate on the "currency_code" field.
func CurrencyCodeHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldCurrencyCode, v))
}

// CurrencyCodeEqualFold applies the EqualFold predicate on the "currency_code" field.
func CurrencyCodeEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldCurrencyCode, v))
}

// CurrencyCodeContainsFold applies the ContainsFold predicate on the "currency_code" field.
func CurrencyCodeContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldCurrencyCode, v))
}

// CardTypeEQ applies the EQ predicate on the "card_type" field.
func CardTypeEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCardType, v))
}

// CardTypeNEQ applies the NEQ predicate on the "card_type" field.
func CardTypeNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCardType, v))
}

// CardTypeIn applies the In predicate on the "card_type" field.
func CardTypeIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCardType, vs...))
}

// CardTypeNotIn applies the NotIn predicate on the "card_type" field.
func CardTypeNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCardType, vs...))
}

// CardTypeGT applies the GT predicate on the "card_type" field.
func CardTypeGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCardType, v))
}

// CardTypeGTE applies the GTE predicate on the "card_type" field.
func CardTypeGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCardType, v))
}

// CardTypeLT applies the LT predicate on the "card_type" field.
func CardTypeLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCardType, v))
}

// CardTypeLTE applies the LTE predicate on the "card_type" field.
func CardTypeLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCardType, v))
}

// CardTypeContains applies the Contains predicate on the "card_type" field.
func CardTypeContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldCardType, v))
}

// CardTypeHasPrefix applies the HasPrefix predicate on the "card_type" field.
func CardTypeHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldCardType, v))
}

// CardTypeHasSuffix applies the HasSuffix predicate on the "card_type" field.
func CardTypeHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldCardType, v))
}

// CardTypeEqualFold applies the EqualFold predicate on the "card_type" field.
func CardTypeEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldCardType, v))
}

// CardTypeContainsFold applies the ContainsFold predicate on the "card_type" field.
func CardTypeContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldCardType, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFinancialAccount applies the HasEdge predicate on the "financial_account" edge.
func HasFinancialAccount() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FinancialAccountTable, FinancialAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFinancialAccountWith applies the HasEdge predicate on the "financial_account" edge with a given conditions (other predicates).
func HasFinancialAccountWith(preds ...predicate.FinancialAccount) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FinancialAccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FinancialAccountTable, FinancialAccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		p(s.Not())
	})
}
