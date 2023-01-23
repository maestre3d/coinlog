package card

import (
	"time"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/domain/user"
)

// Card financial instrument issued by a financial institution as component for interactions between a User and either
// a FinancialAccount or credit line.
type Card struct {
	ID               string
	User             user.User                          // req, FK -> users
	FinancialAccount *financialaccount.FinancialAccount // FK -> financial_accounts (nullable on credit cards)
	DisplayName      string                             // req
	BankName         string
	LastDigits       uint16
	CardType         TypeCard // enum: credit/debit, req
	Balance          float64  // nullable
	BalanceLimit     float64  // nullable
	CurrencyCode     string
	customtype.Auditable
}

func newCard(cmd CreateCommand) (Card, error) {
	if err := domain.Validate.Struct(cmd); err != nil {
		return Card{}, err
	}

	cardType, err := NewTypeCardSafe(cmd.CardType)
	if err != nil {
		return Card{}, err
	}

	if cardType.IsCredit() && cmd.FinancialAccountID != "" {
		return Card{}, ErrCreditWithAccount
	}

	var acc *financialaccount.FinancialAccount
	if cmd.FinancialAccountID != "" {
		acc = &financialaccount.FinancialAccount{
			ID: cmd.FinancialAccountID,
		}
	}

	var balanceLim float64
	if cardType.IsCredit() {
		balanceLim = cmd.BalanceLimit
	}
	return Card{
		ID: cmd.CardID,
		User: user.User{
			ID: cmd.OwnerID,
		},
		FinancialAccount: acc,
		DisplayName:      cmd.DisplayName,
		BankName:         cmd.BankName,
		LastDigits:       uint16(cmd.LastDigits) >> 0,
		CardType:         cardType,
		Balance:          cmd.Balance,
		BalanceLimit:     balanceLim,
		CurrencyCode:     cmd.CurrencyCode,
		Auditable:        customtype.NewAuditable(),
	}, nil
}

func (c *Card) update(cmd UpdateCommand) error {
	if err := domain.Validate.Struct(cmd); err != nil {
		return err
	}

	c.DisplayName = cmd.DisplayName
	c.BankName = cmd.BankName
	c.LastDigits = uint16(cmd.LastDigits) >> 0
	c.Balance = cmd.Balance
	if c.CardType.IsCredit() {
		c.BalanceLimit = cmd.BalanceLimit
	}
	c.CurrencyCode = cmd.CurrencyCode
	c.UpdatedAt = time.Now().UTC()
	c.Version++
	return nil
}
