package exception

import "fmt"

type MissingParameter struct {
	Field string
}

var _ Exception = MissingParameter{}

func (e MissingParameter) Error() string {
	return e.String()
}

func (e MissingParameter) String() string {
	return fmt.Sprintf("Parameter %s is required", e.Field)
}

func (e MissingParameter) TypeName() string {
	return newExceptionTypeName(e)
}
