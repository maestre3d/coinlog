package customtype

import "fmt"

// Enum set of enumerated values.
type Enum interface {
	fmt.Stringer
}
