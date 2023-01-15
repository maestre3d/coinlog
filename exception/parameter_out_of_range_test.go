package exception_test

import (
	"testing"

	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
)

func TestParameterOutOfRange(t *testing.T) {
	tests := []struct {
		name        string
		inField     string
		inFrom      string
		inTo        string
		inLen       string
		expStr      string
		expErr      string
		expTypeName string
	}{
		{
			name:        "empty",
			inField:     "",
			inFrom:      "",
			inTo:        "",
			inLen:       "",
			expStr:      "Parameter  is out of range",
			expErr:      "Parameter  is out of range",
			expTypeName: "ParameterOutOfRange",
		},
		{
			name:        "field only",
			inField:     "foo",
			inFrom:      "",
			inTo:        "",
			inLen:       "",
			expStr:      "Parameter foo is out of range",
			expErr:      "Parameter foo is out of range",
			expTypeName: "ParameterOutOfRange",
		},
		{
			name:        "field and len",
			inField:     "foo",
			inFrom:      "",
			inTo:        "",
			inLen:       "10",
			expStr:      "Parameter foo is out of range [10]",
			expErr:      "Parameter foo is out of range [10]",
			expTypeName: "ParameterOutOfRange",
		},
		{
			name:        "field and from",
			inField:     "foo",
			inFrom:      "10",
			inTo:        "",
			inLen:       "",
			expStr:      "Parameter foo is out of range [10,n]",
			expErr:      "Parameter foo is out of range [10,n]",
			expTypeName: "ParameterOutOfRange",
		},
		{
			name:        "field and to",
			inField:     "foo",
			inFrom:      "",
			inTo:        "10",
			inLen:       "",
			expStr:      "Parameter foo is out of range [0,10]",
			expErr:      "Parameter foo is out of range [0,10]",
			expTypeName: "ParameterOutOfRange",
		},
		{
			name:        "field, from and to",
			inField:     "foo",
			inFrom:      "10",
			inTo:        "100",
			inLen:       "",
			expStr:      "Parameter foo is out of range [10,100]",
			expErr:      "Parameter foo is out of range [10,100]",
			expTypeName: "ParameterOutOfRange",
		},
		{
			name:        "full",
			inField:     "foo",
			inFrom:      "10",
			inTo:        "100",
			inLen:       "99",
			expStr:      "Parameter foo is out of range [10,100]",
			expErr:      "Parameter foo is out of range [10,100]",
			expTypeName: "ParameterOutOfRange",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := exception.ParameterOutOfRange{
				Field: tt.inField,
				From:  tt.inFrom,
				To:    tt.inTo,
				Len:   tt.inLen,
			}
			assert.Equal(t, tt.expStr, out.String())
			assert.Equal(t, tt.expErr, out.Error())
			assert.Equal(t, tt.expTypeName, out.TypeName())
		})
	}
}
