package financialaccount

import "github.com/maestre3d/coinlog/exception"

var (
	ErrNotFound           = exception.ResourceNotFound{Resource: "financial_account"}
	ErrInvalidAccountType = exception.InvalidParameter{
		Field:       "account_type",
		ValidValues: "CHECKING,SAVINGS",
	}
)
