package card

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		name   string
		in     CreateCommand
		exp    Card
		expErr error
	}{
		{
			name:   "empty",
			in:     CreateCommand{},
			exp:    Card{},
			expErr: validator.ValidationErrors{},
		},
		{
			name: "invalid args",
			in: CreateCommand{
				mutateCommand: mutateCommand{
					CardID:       "123",
					DisplayName:  "",
					BankName:     "",
					LastDigits:   0,
					Balance:      0,
					CurrencyCode: "",
				},
				OwnerID:            "",
				FinancialAccountID: "456",
				CardType:           "",
			},
			exp:    Card{},
			expErr: validator.ValidationErrors{},
		},
		{
			name: "missing owner id",
			in: CreateCommand{
				mutateCommand: mutateCommand{
					CardID:       "123",
					DisplayName:  "Foo",
					BankName:     "Bank of America",
					LastDigits:   5675,
					Balance:      45763,
					CurrencyCode: "USD",
				},
				OwnerID:            "",
				FinancialAccountID: "456",
				CardType:           "DEBIT",
			},
			exp:    Card{},
			expErr: validator.ValidationErrors{},
		},
		{
			name: "invalid card type",
			in: CreateCommand{
				mutateCommand: mutateCommand{
					CardID:       "123",
					DisplayName:  "Foo",
					BankName:     "Bank of America",
					LastDigits:   5675,
					Balance:      45763,
					CurrencyCode: "USD",
				},
				OwnerID:            "098",
				FinancialAccountID: "456",
				CardType:           "MISC",
			},
			exp:    Card{},
			expErr: validator.ValidationErrors{},
		},
		{
			name: "invalid credit card with fin_account",
			in: CreateCommand{
				mutateCommand: mutateCommand{
					CardID:       "123",
					DisplayName:  "Foo",
					BankName:     "Bank of America",
					LastDigits:   5675,
					Balance:      45763,
					CurrencyCode: "USD",
				},
				OwnerID:            "098",
				FinancialAccountID: "456",
				CardType:           "CREDIT",
			},
			exp:    Card{},
			expErr: ErrCreditWithAccount,
		},
		{
			name: "valid",
			in: CreateCommand{
				mutateCommand: mutateCommand{
					CardID:       "123",
					DisplayName:  "Foo",
					BankName:     "Bank of America",
					LastDigits:   5675,
					Balance:      45763,
					CurrencyCode: "USD",
				},
				OwnerID:            "098",
				FinancialAccountID: "456",
				CardType:           "DEBIT",
			},
			exp: Card{
				ID: "123",
				User: user.User{
					ID: "098",
				},
				FinancialAccount: &financialaccount.FinancialAccount{
					ID: "456",
				},
				DisplayName:  "Foo",
				BankName:     "Bank of America",
				LastDigits:   5675,
				CardType:     typeCardDebit,
				Balance:      45763,
				CurrencyCode: "USD",
				Auditable: customtype.Auditable{
					IsActive: true,
					Version:  1,
				},
			},
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := newCard(tt.in)
			assert.IsType(t, tt.expErr, err)
			assert.Equal(t, tt.exp.ID, out.ID)
			assert.Equal(t, tt.exp.DisplayName, out.DisplayName)
			assert.Equal(t, tt.exp.BankName, out.BankName)
			assert.Equal(t, tt.exp.LastDigits, out.LastDigits)
			assert.Equal(t, tt.exp.CardType, out.CardType)
			assert.Equal(t, tt.exp.Balance, out.Balance)
			assert.Equal(t, tt.exp.CurrencyCode, out.CurrencyCode)
			assert.Equal(t, tt.exp.Auditable.IsActive, out.Auditable.IsActive)
			assert.Equal(t, tt.exp.Auditable.Version, out.Auditable.Version)

			if tt.exp.FinancialAccount != nil {
				require.NotNil(t, out.FinancialAccount)
				assert.Equal(t, tt.exp.FinancialAccount.ID, out.FinancialAccount.ID)
			}
		})
	}
}
