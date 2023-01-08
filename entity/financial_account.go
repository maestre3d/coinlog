package entity

import "github.com/maestre3d/coinlog/valueobject"

// FinancialAccount storage for financial assets of a User backed and held by a financial institution.
type FinancialAccount struct {
	ID              string
	User            User   // req, FK -> users
	DisplayName     string // req
	InstitutionName string
	AccountType     string // enum: checking/savings/...
	Balance         float64
	Auditable       valueobject.Auditable
}
