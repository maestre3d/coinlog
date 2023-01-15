package exception

import "fmt"

type ResourceNotFound struct {
	Resource string
}

var _ Exception = ResourceNotFound{}

func (e ResourceNotFound) Error() string {
	return e.String()
}

func (e ResourceNotFound) String() string {
	return fmt.Sprintf("Resource %s not found", e.Resource)
}

func (e ResourceNotFound) TypeName() string {
	return newExceptionTypeName(e)
}
