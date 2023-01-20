package financialaccount

type mutateCommand struct {
	AccountID    string  `json:"account_id" validate:"required"`
	DisplayName  string  `json:"display_name" validate:"required,lte=36"`
	BankName     string  `json:"bank_name" validate:"omitempty,lte=64"`
	AccountType  string  `json:"account_type" validate:"required,oneof=CHECKING SAVINGS"`
	Balance      float64 `json:"balance" validate:"omitempty,gte=-999999999999,lte=999999999999"`
	CurrencyCode string  `json:"currency_code" validate:"required,uppercase,iso4217"`
}

type CreateCommand struct {
	mutateCommand
	OwnerID string `json:"owner_id" validate:"required"`
}

type UpdateCommand struct {
	mutateCommand
}
