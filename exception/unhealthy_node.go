package exception

type UnhealthyNode struct {
	Parent error
}

var _ Exception = UnhealthyNode{}

var _ Wrapper = UnhealthyNode{}

func (u UnhealthyNode) TypeName() string {
	return "UnhealthyNode"
}

func (u UnhealthyNode) Error() string {
	return u.Parent.Error()
}

func (u UnhealthyNode) String() string {
	return u.Parent.Error()
}

func (u UnhealthyNode) Unwrap() error {
	return u.Parent
}
