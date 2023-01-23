package card

import (
	"errors"

	"github.com/maestre3d/coinlog/exception"
)

var (
	ErrNotFound        = exception.ResourceNotFound{Resource: "card"}
	ErrInvalidCardType = exception.InvalidParameter{
		Field:       "card",
		ValidValues: "CREDIT,DEBIT",
	}
	ErrCreditWithAccount = exception.DomainGeneric{
		Code:   "InvalidCreditCard",
		Parent: errors.New("credit cards are not allowed to attach financial accounts"),
	}
)
