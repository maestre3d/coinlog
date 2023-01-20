package exception_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFromValidator(t *testing.T) {
	validate := validator.New()
	type fakeA struct {
		Req    string `validate:"required"`
		Len    string `validate:"len=3"`
		Multi  string `validate:"required,gte=4"`
		FooBar string `validate:"gte=4"`
	}

	tests := []struct {
		name        string
		in          interface{}
		expTypeName []string
		expMsg      []string
	}{
		{
			name:        "empty",
			in:          fakeA{},
			expTypeName: []string{"MissingParameter", "ParameterOutOfRange", "MissingParameter", "ParameterOutOfRange"},
			expMsg: []string{
				"Parameter req is required",
				"Parameter len is out of range [3]",
				"Parameter multi is required",
				"Parameter foo_bar is out of range [4,n]",
			},
		},
		{
			name: "multi second err",
			in: fakeA{
				Multi: "foo",
			},
			expTypeName: []string{"MissingParameter", "ParameterOutOfRange", "ParameterOutOfRange", "ParameterOutOfRange"},
		},
		{
			name: "last item err",
			in: fakeA{
				Req:    "f",
				Len:    "foo",
				Multi:  "foobar",
				FooBar: "foo",
			},
			expTypeName: []string{"ParameterOutOfRange"},
		},
		{
			name: "invalid param oneof",
			in: struct {
				A string `validate:"oneof=foo bar"`
			}{
				A: "baz",
			},
			expTypeName: []string{"InvalidParameter"},
		},
		{
			name: "low param",
			in: struct {
				A string `validate:"min=4"`
			}{
				A: "foo",
			},
			expTypeName: []string{"ParameterOutOfRange"},
		},
		{
			name: "invalid param ipv4",
			in: struct {
				A string `validate:"ipv4"`
			}{
				A: "10.8.FOO.1",
			},
			expTypeName: []string{"InvalidParameter"},
		},
		{
			name: "unsupported type",
			in: struct {
				A string `validate:"bic"`
			}{
				A: "dummy",
			},
			expTypeName: []string{"DomainGeneric"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.in)
			require.NotNil(t, err)
			errs, okErrs := err.(validator.ValidationErrors)
			require.True(t, okErrs)
			require.Len(t, errs, len(tt.expTypeName))
			for i, e := range errs {
				errValid, ok := e.(validator.FieldError)
				require.True(t, ok)
				ex := exception.NewFromValidator(errValid)
				assert.Equal(t, tt.expTypeName[i], ex.TypeName())
				if len(tt.expMsg) > 0 {
					assert.Equal(t, tt.expMsg[i], ex.String())
				}
			}
		})
	}
}
