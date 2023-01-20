package customtype_test

import (
	"testing"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/stretchr/testify/assert"
)

func TestNewAuditable(t *testing.T) {
	out := customtype.NewAuditable()
	assert.Equal(t, uint32(1), out.Version)
	assert.True(t, out.IsActive)
	assert.NotZero(t, out.CreatedAt)
	assert.EqualValues(t, out.CreatedAt, out.UpdatedAt)
}
