package valueobject_test

import (
	"testing"

	"github.com/maestre3d/coinlog/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewPageToken(t *testing.T) {
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
			name: "digit",
			in:   "1",
			exp:  "MQ==",
		},
		{
			name: "numeric",
			in:   "123",
			exp:  "MTIz",
		},
		{
			name: "alphabetical",
			in:   "abc",
			exp:  "YWJj",
		},
		{
			name: "alphanumerical",
			in:   "123abc",
			exp:  "MTIzYWJj",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := valueobject.NewPageToken(tt.in)
			assert.Equal(t, tt.exp, out.String())
			assert.Equal(t, tt.in, valueobject.DecodePageToken(out))
		})
	}
}
