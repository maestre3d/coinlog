package exception

type DomainGeneric struct {
	Parent error
}

var _ Exception = DomainGeneric{}

func (e DomainGeneric) Error() string {
	return e.Parent.Error()
}

func (e DomainGeneric) String() string {
	return e.Parent.Error()
}

func (e DomainGeneric) TypeName() string {
	return newExceptionTypeName(e)
}