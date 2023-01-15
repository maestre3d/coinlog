package exception

import "fmt"

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
	case e.From != "" && e.To != "":
		return fmt.Sprintf("Parameter %s is out of range [%s,%s]", e.Field, e.From, e.To)
	case e.Len != "":
		return fmt.Sprintf("Parameter %s is out of range [%s]", e.Field, e.Len)
	case e.From != "":
		return fmt.Sprintf("Parameter %s is out of range [%s,n]", e.Field, e.From)
	case e.To != "":
		return fmt.Sprintf("Parameter %s is out of range [0,%s]", e.Field, e.To)
	default:
		return fmt.Sprintf("Parameter %s is out of range", e.Field)
	}
}

func (e ParameterOutOfRange) TypeName() string {
	return newExceptionTypeName(e)
}
