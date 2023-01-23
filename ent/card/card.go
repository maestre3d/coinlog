// Code generated by ent, DO NOT EDIT.

package card

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFinancialAccountID holds the string denoting the financial_account_id field in the database.
	FieldFinancialAccountID = "financial_account_id"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldBankName holds the string denoting the bank_name field in the database.
	FieldBankName = "bank_name"
	// FieldLastDigits holds the string denoting the last_digits field in the database.
	FieldLastDigits = "last_digits"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldBalanceLimit holds the string denoting the balance_limit field in the database.
	FieldBalanceLimit = "balance_limit"
	// FieldCurrencyCode holds the string denoting the currency_code field in the database.
	FieldCurrencyCode = "currency_code"
	// FieldCardType holds the string denoting the card_type field in the database.
	FieldCardType = "card_type"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeFinancialAccount holds the string denoting the financial_account edge name in mutations.
	EdgeFinancialAccount = "financial_account"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "cards"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_cards"
	// FinancialAccountTable is the table that holds the financial_account relation/edge.
	FinancialAccountTable = "cards"
	// FinancialAccountInverseTable is the table name for the FinancialAccount entity.
	// It exists in this package in order to avoid circular dependency with the "financialaccount" package.
	FinancialAccountInverseTable = "financial_accounts"
	// FinancialAccountColumn is the table column denoting the financial_account relation/edge.
	FinancialAccountColumn = "financial_account_id"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldIsActive,
	FieldVersion,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFinancialAccountID,
	FieldDisplayName,
	FieldBankName,
	FieldLastDigits,
	FieldBalance,
	FieldBalanceLimit,
	FieldCurrencyCode,
	FieldCardType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_cards",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// CurrencyCodeValidator is a validator for the "currency_code" field. It is called by the builders before save.
	CurrencyCodeValidator func(string) error
	// CardTypeValidator is a validator for the "card_type" field. It is called by the builders before save.
	CardTypeValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
