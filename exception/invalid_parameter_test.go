package exception_test

import (
	"testing"

	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
)

func TestInvalidParameter(t *testing.T) {
	tests := []struct {
		name          string
		inField       string
		inValidValues string
		expStr        string
		expErr        string
		expTypeName   string
	}{
		{
			name:          "empty",
			inField:       "",
			inValidValues: "",
			expStr:        "Parameter  has an invalid format",
			expErr:        "Parameter  has an invalid format",
			expTypeName:   "InvalidParameter",
		},
		{
			name:          "no valid values",
			inField:       "foo",
			inValidValues: "",
			expStr:        "Parameter foo has an invalid format",
			expErr:        "Parameter foo has an invalid format",
			expTypeName:   "InvalidParameter",
		},
		{
			name:          "full data",
			inField:       "foo",
			inValidValues: "bar,baz",
			expStr:        "Parameter foo has an invalid format, expected one of [bar,baz]",
			expErr:        "Parameter foo has an invalid format, expected one of [bar,baz]",
			expTypeName:   "InvalidParameter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := exception.InvalidParameter{
				Field:       tt.inField,
				ValidValues: tt.inValidValues,
			}
			assert.Equal(t, tt.expStr, out.String())
			assert.Equal(t, tt.expErr, out.Error())
			assert.Equal(t, tt.expTypeName, out.TypeName())
		})
	}
}
