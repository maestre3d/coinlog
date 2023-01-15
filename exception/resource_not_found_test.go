package exception_test

import (
	"testing"

	"github.com/maestre3d/coinlog/exception"
	"github.com/stretchr/testify/assert"
)

func TestResourceNotFound(t *testing.T) {
	tests := []struct {
		name        string
		inRsc       string
		expStr      string
		expErr      string
		expTypeName string
	}{
		{
			name:        "empty",
			inRsc:       "",
			expStr:      "Resource  not found",
			expErr:      "Resource  not found",
			expTypeName: "ResourceNotFound",
		},
		{
			name:        "full data",
			inRsc:       "foo",
			expStr:      "Resource foo not found",
			expErr:      "Resource foo not found",
			expTypeName: "ResourceNotFound",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := exception.ResourceNotFound{
				Resource: tt.inRsc,
			}
			assert.Equal(t, tt.expStr, out.String())
			assert.Equal(t, tt.expErr, out.Error())
			assert.Equal(t, tt.expTypeName, out.TypeName())
		})
	}
}
