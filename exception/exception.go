// Package exception system-related errors.
package exception

import "fmt"

// Exception system-related error container.
type Exception interface {
	error
	fmt.Stringer
	TypeName() string
}
