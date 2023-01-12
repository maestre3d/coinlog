package card

import (
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/domain/user"
)

// Card financial instrument issued by a financial institution as component for interactions between a User and either
// a FinancialAccount or credit line.
type Card struct {
	ID               string
	User             user.User                         // req, FK -> users
	FinancialAccount financialaccount.FinancialAccount // FK -> financial_accounts (nullable on credit cards)
	DisplayName      string                            // req
	InstitutionName  string
	LastDigits       uint8
	CardType         TypeCard // enum: credit/debit, req
	Balance          float64  // nullable
	domain.Auditable
}

type NewCardArgs struct {
	ID                 string
	UserID             string
	FinancialAccountID string
	DisplayName        string
	InstitutionName    string
	LastDigits         int
}

func NewCard() Card {
	a := Card{}
	return a
}
