package exception_test

import (
	"testing"

	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
)

func TestMissingParameter(t *testing.T) {
	tests := []struct {
		name        string
		inField     string
		expStr      string
		expErr      string
		expTypeName string
	}{
		{
			name:        "empty",
			inField:     "",
			expStr:      "Parameter  is required",
			expErr:      "Parameter  is required",
			expTypeName: "MissingParameter",
		},
		{
			name:        "full data",
			inField:     "foo",
			expStr:      "Parameter foo is required",
			expErr:      "Parameter foo is required",
			expTypeName: "MissingParameter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := exception.MissingParameter{
				Field: tt.inField,
			}
			assert.Equal(t, tt.expStr, out.String())
			assert.Equal(t, tt.expErr, out.Error())
			assert.Equal(t, tt.expTypeName, out.TypeName())
		})
	}
}
