package card

type mutateCommand struct {
	CardID       string  `json:"account_id" validate:"required"`
	DisplayName  string  `json:"display_name" validate:"required,lte=36"`
	BankName     string  `json:"bank_name" validate:"omitempty,lte=64"`
	LastDigits   int     `json:"last_digits" validate:"omitempty,lte=9999"`
	Balance      float64 `json:"balance" validate:"omitempty,gte=-999999999999,lte=999999999999"`
	BalanceLimit float64 `json:"balance_limit" validate:"omitempty,gte=-999999999999,lte=999999999999"`
	CurrencyCode string  `json:"currency_code" validate:"required,uppercase,iso4217"`
}

type CreateCommand struct {
	mutateCommand
	OwnerID            string `json:"owner_id" validate:"required"`
	FinancialAccountID string `json:"financial_account_id"`
	CardType           string `json:"card_type" validate:"required,oneof=CREDIT DEBIT"`
}

type UpdateCommand struct {
	mutateCommand
}
