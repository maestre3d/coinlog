package domain_test

import (
	"testing"

	"github.com/maestre3d/coinlog/domain"
)

type noPtrNullable struct {
	foo string
}

var _ domain.Nullable[noPtrNullable] = noPtrNullable{}

func (n noPtrNullable) PtrIfNotEmpty() *noPtrNullable {
	if n.foo == "" {
		return nil
	}

	return &n
}

type ptrNullable struct {
	foo string
}

var _ domain.Nullable[ptrNullable] = &ptrNullable{}

func (p *ptrNullable) PtrIfNotEmpty() *ptrNullable {
	if p.foo == "" {
		return nil
	}
	return p
}

func TestNullable(t *testing.T) {
	noPtr := noPtrNullable{}
	t.Log(noPtr.PtrIfNotEmpty())
	noPtr.foo = "foo"
	t.Log(noPtr.PtrIfNotEmpty())

	ptrVal := &ptrNullable{}
	t.Log(ptrVal.PtrIfNotEmpty())
	ptrVal.foo = "foo"
	t.Log(ptrVal.PtrIfNotEmpty())
}
