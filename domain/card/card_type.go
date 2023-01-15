package card

import (
	"github.com/maestre3d/coinlog/customtype"
)

const (
	typeCardUnknown int = iota
	typeCardCredit
	typeCardDebit
)

var typeCardMap = map[int]string{
	typeCardUnknown: "UNKNOWN",
	typeCardCredit:  "CREDIT",
	typeCardDebit:   "DEBIT",
}

type TypeCard int

var _ customtype.Enum = TypeCard(1)

func (c TypeCard) String() string {
	return typeCardMap[int(c)]
}
