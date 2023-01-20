package exception_test

import (
	"errors"
	"testing"

	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
)

func TestNewUnhealthyNode(t *testing.T) {
	var errPsql = errors.New("psql error")
	ex := exception.UnhealthyNode{
		Parent: errPsql,
	}
	assert.Equal(t, "psql error", ex.Error())
	assert.Equal(t, "psql error", ex.String())
	assert.Equal(t, "UnhealthyNode", ex.TypeName())
	assert.EqualValues(t, errPsql, ex.Unwrap())
	assert.Equal(t, errPsql, errors.Unwrap(ex))
}
