package exception_test

import (
	"errors"
	"testing"

	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainGeneric(t *testing.T) {
	var fakeErr = errors.New("foo error")
	ex := exception.DomainGeneric{
		Code:   "MissingFoo",
		Parent: fakeErr,
	}
	assert.Equal(t, "foo error", ex.Error())
	assert.Equal(t, "foo error", ex.String())
	assert.Equal(t, "MissingFoo", ex.TypeName())
	assert.EqualValues(t, fakeErr, ex.Unwrap())
	assert.Equal(t, fakeErr, errors.Unwrap(ex))
}
