package entity

import (
	"time"

	"github.com/maestre3d/coinlog/valueobject"
)

// FinancialTransaction record generated by the financial interaction between a User and a Contact.
//
// Assets could come from either FinancialAccount directly or from a Card; both instruments are mutually exclusive.
//
// Transactions could also be generated by ScheduledPayment(s).
type FinancialTransaction struct {
	ID                   string
	User                 User             // req, FK -> users
	FinancialAccount     FinancialAccount // FK -> financial_accounts, nullable
	Card                 Card             // FK -> cards, nullable
	Contact              Contact          // FK -> contacts
	FromScheduledPayment ScheduledPayment // FK -> scheduled_payments, nullable
	DisplayName          string           // req
	Description          string
	Amount               float64
	Currency             string    // uses ISO-4127 standard
	OccurredAt           time.Time //req
	Auditable            valueobject.Auditable
}