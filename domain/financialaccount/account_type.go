package financialaccount

import (
	"github.com/maestre3d/coinlog/customtype"
)

type AccountType uint8

const (
	accountTypeUnknown AccountType = iota
	accountTypeChecking
	accountTypeSavings
)

var _ customtype.Enum = AccountType(1)

var accountTypeEnumMap = map[AccountType]string{
	accountTypeUnknown:  "UNKNOWN",
	accountTypeChecking: "CHECKING",
	accountTypeSavings:  "SAVINGS",
}

var accountTypeEnumStringMap = map[string]AccountType{
	"UNKNOWN":  accountTypeUnknown,
	"CHECKING": accountTypeChecking,
	"SAVINGS":  accountTypeSavings,
}

func NewAccountType(v string) AccountType {
	return accountTypeEnumStringMap[v]
}

func NewAccountTypeSafe(v string) (AccountType, error) {
	out, ok := accountTypeEnumStringMap[v]
	if !ok {
		return 0, ErrInvalidAccountType
	}
	return out, nil
}

func (a AccountType) String() string {
	return accountTypeEnumMap[a]
}
