package financialaccount

import (
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/user"
)

// FinancialAccount storage for financial assets of a User backed and held by a financial institution.
type FinancialAccount struct {
	ID              string
	User            user.User // req, FK -> users
	DisplayName     string    // req
	InstitutionName string
	AccountType     AccountType // enum: checking/savings/...
	Balance         float64
	domain.Auditable
}
