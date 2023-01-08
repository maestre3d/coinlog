package exception

import (
	"strings"

	"github.com/modern-go/reflect2"
)

func newExceptionTypeName(v any) string {
	return strings.TrimPrefix(reflect2.TypeOf(v).String(), "exception.")
}
