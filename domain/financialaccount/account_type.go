package financialaccount

import (
	"github.com/maestre3d/coinlog/customtype"
)

const (
	accountTypeUnknown int = iota
	accountTypeChecking
	accountTypeSavings
)

var accountTypeEnumMap = map[int]string{
	accountTypeUnknown:  "UNKNOWN",
	accountTypeChecking: "CHECKING",
	accountTypeSavings:  "SAVINGS",
}

type AccountType int

var _ customtype.Enum = AccountType(1)

func (a AccountType) String() string {
	return accountTypeEnumMap[int(a)]
}
