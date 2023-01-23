package card

import (
	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/pointer"
)

type View struct {
	ID               string                 `json:"card_id"`
	User             *user.View             `json:"user"`
	FinancialAccount *financialaccount.View `json:"financial_account"`
	DisplayName      string                 `json:"display_name"`
	BankName         *string                `json:"bank_name"`
	LastDigits       *uint16                `json:"last_digits"`
	CardType         string                 `json:"card_type"`
	Balance          float64                `json:"balance"`
	BalanceLimit     *float64               `json:"balance_limit"`
	CurrencyCode     string                 `json:"currency_code"`
	customtype.Auditable
}

func NewView(c Card) View {
	var acc *financialaccount.View
	if c.FinancialAccount != nil {
		acc = pointer.PtrTo(financialaccount.NewView(*c.FinancialAccount))
	}
	return View{
		ID:               c.ID,
		User:             pointer.PtrIfNotEmpty(user.NewView(c.User)),
		FinancialAccount: acc,
		DisplayName:      c.DisplayName,
		BankName:         pointer.PtrIfNotEmpty(c.BankName),
		LastDigits:       pointer.PtrIfNotEmpty(c.LastDigits),
		CardType:         c.CardType.String(),
		Balance:          c.Balance,
		BalanceLimit:     pointer.PtrIfNotEmpty(c.BalanceLimit),
		CurrencyCode:     c.CurrencyCode,
		Auditable: customtype.Auditable{
			IsActive:  c.IsActive,
			Version:   c.Version,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
	}
}
