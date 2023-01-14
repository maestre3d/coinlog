package domain_test

import (
	"testing"

	"github.com/maestre3d/coinlog/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtrIfNotEmpty(t *testing.T) {
	// string
	s := domain.PtrIfNotEmpty("")
	assert.Nil(t, s)
	s = domain.PtrIfNotEmpty("foo")
	assert.EqualValues(t, domain.PtrTo("foo"), s)

	// int
	i := domain.PtrIfNotEmpty(0)
	assert.Nil(t, i)
	i = domain.PtrIfNotEmpty(10)
	assert.EqualValues(t, domain.PtrTo(10), i)
}

func TestPtrTo(t *testing.T) {
	tests := []struct {
		name string
		in   any
		exp  any
	}{
		{
			name: "empty string",
			in:   "",
			exp:  domain.PtrTo(""),
		},
		{
			name: "empty int",
			in:   0,
			exp:  domain.PtrTo(0),
		},
		{
			name: "string",
			in:   "foo",
			exp:  domain.PtrTo("foo"),
		},
		{
			name: "int",
			in:   10,
			exp:  domain.PtrTo(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := domain.PtrTo(tt.in)
			require.NotNil(t, out)
			switch v := tt.exp.(type) {
			case *string:
				assert.EqualValues(t, *v, *out)
			case *int:
				assert.EqualValues(t, *v, *out)
			}
		})
	}
}
