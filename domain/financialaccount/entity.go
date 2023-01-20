package financialaccount

import (
	"time"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/user"
)

// FinancialAccount storage for financial assets of a User backed and held by a financial institution.
type FinancialAccount struct {
	ID           string
	User         user.User // req, FK -> users
	DisplayName  string    // req
	BankName     string
	AccountType  AccountType // enum: checking/savings/...
	Balance      float64
	CurrencyCode string
	customtype.Auditable
}

func newFinancialAccount(cmd CreateCommand) (FinancialAccount, error) {
	if err := domain.Validate.Struct(cmd); err != nil {
		return FinancialAccount{}, err
	}

	accType, err := NewAccountTypeSafe(cmd.AccountType)
	if err != nil {
		return FinancialAccount{}, err
	}

	return FinancialAccount{
		ID: cmd.AccountID,
		User: user.User{
			ID: cmd.OwnerID,
		},
		DisplayName:  cmd.DisplayName,
		BankName:     cmd.BankName,
		AccountType:  accType,
		Balance:      cmd.Balance,
		CurrencyCode: cmd.CurrencyCode,
		Auditable:    customtype.NewAuditable(),
	}, nil
}

func (f *FinancialAccount) update(cmd UpdateCommand) error {
	if err := domain.Validate.Struct(cmd); err != nil {
		return err
	}

	f.DisplayName = cmd.DisplayName
	f.BankName = cmd.BankName
	f.AccountType = NewAccountType(cmd.AccountType)
	f.Balance = cmd.Balance
	f.Version++
	f.UpdatedAt = time.Now().UTC()
	return nil
}
