package scheduledpayment

import (
	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/card"
	"github.com/maestre3d/coinlog/domain/contact"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/domain/user"
)

// ScheduledPayment interaction where an external agent (Contact) periodically requests financial assets from a
// User's FinancialAccount or Card.
//
// Assets could come from either FinancialAccount directly or from a Card; both instruments are mutually exclusive.
type ScheduledPayment struct {
	ID               string
	User             user.User                          // FK -> users
	FinancialAccount *financialaccount.FinancialAccount // FK -> financial_accounts, nullable
	Card             *card.Card                         // FK -> cards, nullable
	Contact          contact.Contact                    // FK -> contacts
	DisplayName      string                             // req
	Description      string
	Amount           float64
	CurrencyCode     string // uses ISO-4127 standard
	IntervalDays     uint16 // req
	customtype.Auditable
}
