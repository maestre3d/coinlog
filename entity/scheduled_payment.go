package entity

import "github.com/maestre3d/coinlog/valueobject"

// ScheduledPayment interaction where an external agent (Contact) periodically requests financial assets from a
// User's FinancialAccount or Card.
//
// Assets could come from either FinancialAccount directly or from a Card; both instruments are mutually exclusive.
type ScheduledPayment struct {
	ID               string
	User             User             // FK -> users
	FinancialAccount FinancialAccount // FK -> financial_accounts, nullable
	Card             Card             // FK -> cards, nullable
	Contact          Contact          // FK -> contacts
	DisplayName      string           // req
	Description      string
	Amount           float64
	Currency         string // uses ISO-4127 standard
	IntervalDays     uint16 // req
	Auditable        valueobject.Auditable
}
