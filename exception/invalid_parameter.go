package exception

import "fmt"

type InvalidParameter struct {
	Field       string
	ValidValues string
}

var _ Exception = InvalidParameter{}

func (e InvalidParameter) Error() string {
	return e.String()
}

func (e InvalidParameter) String() string {
	if e.ValidValues != "" {
		return fmt.Sprintf("Parameter %s has an invalid format, expected one of [%s]", e.Field, e.ValidValues)
	}
	return fmt.Sprintf("Parameter %s has an invalid format", e.Field)
}

func (e InvalidParameter) TypeName() string {
	return newExceptionTypeName(e)
}
