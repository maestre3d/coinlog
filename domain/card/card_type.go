package card

import (
	"github.com/maestre3d/coinlog/customtype"
)

const (
	typeCardUnknown TypeCard = iota
	typeCardCredit
	typeCardDebit
)

var typeCardMap = map[TypeCard]string{
	typeCardUnknown: "UNKNOWN",
	typeCardCredit:  "CREDIT",
	typeCardDebit:   "DEBIT",
}

var typeCardStringMap = map[string]TypeCard{
	"UNKNOWN": typeCardUnknown,
	"CREDIT":  typeCardCredit,
	"DEBIT":   typeCardDebit,
}

type TypeCard uint8

var _ customtype.Enum = TypeCard(1)

func NewTypeCard(v string) TypeCard {
	return typeCardStringMap[v]
}

func NewTypeCardSafe(v string) (TypeCard, error) {
	tp, ok := typeCardStringMap[v]
	if !ok {
		return 0, ErrInvalidCardType
	}
	return tp, nil
}

func (c TypeCard) String() string {
	return typeCardMap[c]
}

func (c TypeCard) IsCredit() bool {
	return c == typeCardCredit
}

func (c TypeCard) IsDebit() bool {
	return c == typeCardDebit
}
