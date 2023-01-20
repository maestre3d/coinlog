// Package exception system-related errors.
package exception

import "fmt"

// Exception system-related error container.
type Exception interface {
	error
	fmt.Stringer
	TypeName() string
}

// Wrapper custom Exception type with parent error capabilities.
type Wrapper interface {
	// Unwrap retrieves parent error.
	Unwrap() error
}
