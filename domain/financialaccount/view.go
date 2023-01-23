package financialaccount

import (
	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/pointer"
)

type View struct {
	ID           string     `json:"financial_account_id"`
	User         *user.View `json:"user"`
	DisplayName  string     `json:"display_name"`
	BankName     *string    `json:"bank_name"`
	AccountType  string     `json:"account_type"`
	Balance      float64    `json:"balance"`
	CurrencyCode string     `json:"currency_code"`
	customtype.Auditable
}

func NewView(acc FinancialAccount) View {
	return View{
		ID:           acc.ID,
		User:         pointer.PtrIfNotEmpty(user.NewView(acc.User)),
		DisplayName:  acc.DisplayName,
		BankName:     pointer.PtrIfNotEmpty(acc.BankName),
		AccountType:  acc.AccountType.String(),
		Balance:      acc.Balance,
		CurrencyCode: acc.CurrencyCode,
		Auditable: customtype.Auditable{
			IsActive:  acc.IsActive,
			Version:   acc.Version,
			CreatedAt: acc.CreatedAt,
			UpdatedAt: acc.UpdatedAt,
		},
	}
}
