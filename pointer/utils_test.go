package pointer_test

import (
	"testing"

	"github.com/maestre3d/coinlog/pointer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtrIfNotEmpty(t *testing.T) {
	// string
	s := pointer.PtrIfNotEmpty("")
	assert.Nil(t, s)
	s = pointer.PtrIfNotEmpty("foo")
	assert.EqualValues(t, pointer.PtrTo("foo"), s)

	// int
	i := pointer.PtrIfNotEmpty(0)
	assert.Nil(t, i)
	i = pointer.PtrIfNotEmpty(10)
	assert.EqualValues(t, pointer.PtrTo(10), i)
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
			exp:  pointer.PtrTo(""),
		},
		{
			name: "empty int",
			in:   0,
			exp:  pointer.PtrTo(0),
		},
		{
			name: "string",
			in:   "foo",
			exp:  pointer.PtrTo("foo"),
		},
		{
			name: "int",
			in:   10,
			exp:  pointer.PtrTo(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := pointer.PtrTo(tt.in)
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
