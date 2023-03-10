// Code generated by ent, DO NOT EDIT.

package financialaccount

const (
	// Label holds the string label denoting the financialaccount type in the database.
	Label = "financial_account"
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
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldBankName holds the string denoting the bank_name field in the database.
	FieldBankName = "bank_name"
	// FieldAccountType holds the string denoting the account_type field in the database.
	FieldAccountType = "account_type"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldCurrencyCode holds the string denoting the currency_code field in the database.
	FieldCurrencyCode = "currency_code"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// Table holds the table name of the financialaccount in the database.
	Table = "financial_accounts"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "financial_accounts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_financial_accounts"
	// CardsTable is the table that holds the cards relation/edge.
	CardsTable = "cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// CardsColumn is the table column denoting the cards relation/edge.
	CardsColumn = "financial_account_id"
)

// Columns holds all SQL columns for financialaccount fields.
var Columns = []string{
	FieldID,
	FieldIsActive,
	FieldVersion,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDisplayName,
	FieldBankName,
	FieldAccountType,
	FieldBalance,
	FieldCurrencyCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "financial_accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_financial_accounts",
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
	// AccountTypeValidator is a validator for the "account_type" field. It is called by the builders before save.
	AccountTypeValidator func(string) error
	// CurrencyCodeValidator is a validator for the "currency_code" field. It is called by the builders before save.
	CurrencyCodeValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
