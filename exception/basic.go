package exception

import "fmt"

type DomainUnknown struct {
	Parent error
}

var _ Exception = DomainUnknown{}

func (e DomainUnknown) Error() string {
	return e.String()
}

func (e DomainUnknown) String() string {
	return e.Parent.Error()
}

func (e DomainUnknown) TypeName() string {
	return newExceptionTypeName(e)
}

type ResourceNotFound struct {
	Resource string
}

var _ Exception = ResourceNotFound{}

func (e ResourceNotFound) Error() string {
	return e.String()
}

func (e ResourceNotFound) String() string {
	return fmt.Sprintf("%s was not found", e.Resource)
}

func (e ResourceNotFound) TypeName() string {
	return newExceptionTypeName(e)
}

type MissingParameter struct {
	Field string
}

var _ Exception = MissingParameter{}

func (e MissingParameter) Error() string {
	return e.String()
}

func (e MissingParameter) String() string {
	return fmt.Sprintf("%s is required", e.Field)
}

func (e MissingParameter) TypeName() string {
	return newExceptionTypeName(e)
}

type ParameterOutOfRange struct {
	Field string
	From  string
	To    string
	Len   string
}

var _ Exception = ParameterOutOfRange{}

func (e ParameterOutOfRange) Error() string {
	return e.String()
}

func (e ParameterOutOfRange) String() string {
	switch {
	case e.Len != "":
		return fmt.Sprintf("%s is out of range (%s)", e.Field, e.Len)
	case e.From != "":
		return fmt.Sprintf("%s is out of range [%s,n)", e.Field, e.From)
	case e.To != "":
		return fmt.Sprintf("%s is out of range [0,%s)", e.Field, e.To)
	default:
		return fmt.Sprintf("%s is out of range", e.Field)
	}
}

func (e ParameterOutOfRange) TypeName() string {
	return newExceptionTypeName(e)
}

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
		return fmt.Sprintf("%s is invalid (%s)", e.Field, e.ValidValues)
	}
	return fmt.Sprintf("%s has an invalid format", e.Field)
}

func (e InvalidParameter) TypeName() string {
	return newExceptionTypeName(e)
}
