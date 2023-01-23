package exception

type DomainGeneric struct {
	Code   string
	Parent error
}

var _ Exception = DomainGeneric{}

var _ Wrapper = DomainGeneric{}

func (e DomainGeneric) Error() string {
	return e.Parent.Error()
}

func (e DomainGeneric) String() string {
	return e.Parent.Error()
}

func (e DomainGeneric) TypeName() string {
	return e.Code
}

func (e DomainGeneric) Unwrap() error {
	return e.Parent
}
