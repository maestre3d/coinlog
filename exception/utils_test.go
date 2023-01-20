package exception

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExceptionTypeName(t *testing.T) {
	type exampleFoo struct {
		foo string
	}

	tests := []struct {
		name string
		in   any
		exp  string
	}{
		{
			name: "empty",
			in:   nil,
			exp:  "",
		},
		{
			name: "primitive type",
			in:   "foo",
			exp:  "string",
		},
		{
			name: "custom type",
			in:   exampleFoo{},
			exp:  "exampleFoo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := newExceptionTypeName(tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestNewSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		exp  string
	}{
		{
			name: "empty",
			in:   "",
			exp:  "",
		},
		{
			name: "single",
			in:   "f",
			exp:  "f",
		},
		{
			name: "lower",
			in:   "foobar",
			exp:  "foobar",
		},
		{
			name: "pascal case",
			in:   "FooBar",
			exp:  "foo_bar",
		},
		{
			name: "camel case",
			in:   "fooBar",
			exp:  "foo_bar",
		},
		{
			name: "snake case",
			in:   "foo_bar",
			exp:  "foo_bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := newSnakeCase(tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func BenchmarkNewSnakeCase(b *testing.B) {
	in := "FooBarBaz"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = newSnakeCase(in)
	}
}
