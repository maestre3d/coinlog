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
		Req   string `json:"foo" validate:"required"`
		Len   string `json:"len" validate:"len=3"`
		Multi string `json:"multi" validate:"required,gte=4"`
	}

	type fakeB struct {
		Invalid string `json:"invalid" validate:"oneof=foo bar"`
	}

	type fakeC struct {
		Bic string `json:"bic" validate:"bic"`
	}

	tests := []struct {
		name        string
		in          interface{}
		expTypeName []string
	}{
		{
			name:        "empty",
			in:          fakeA{},
			expTypeName: []string{"MissingParameter", "ParameterOutOfRange", "MissingParameter"},
		},
		{
			name: "multi second err",
			in: fakeA{
				Multi: "foo",
			},
			expTypeName: []string{"MissingParameter", "ParameterOutOfRange", "ParameterOutOfRange"},
		},
		{
			name: "last item err",
			in: fakeA{
				Req:   "f",
				Len:   "foo",
				Multi: "foo",
			},
			expTypeName: []string{"ParameterOutOfRange"},
		},
		{
			name: "invalid param",
			in: fakeB{
				Invalid: "baz",
			},
			expTypeName: []string{"InvalidParameter"},
		},
		{
			name: "unsupported type",
			in: fakeC{
				Bic: "dummy",
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
			}
		})
	}
}
