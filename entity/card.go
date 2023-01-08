package entity

import "github.com/maestre3d/coinlog/valueobject"

// Card financial instrument issued by a financial institution as component for interactions between a User and either
// a FinancialAccount or credit line.
type Card struct {
	ID               string
	User             User             // req, FK -> users
	FinancialAccount FinancialAccount // FK -> financial_accounts (nullable on credit cards)
	DisplayName      string           // req
	InstitutionName  string
	LastDigits       uint8
	CardType         string  // enum: credit/debit, req
	Balance          float64 // nullable
	Auditable        valueobject.Auditable
}
