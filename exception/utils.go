package exception

import (
	"strings"
	"unicode"

	"github.com/modern-go/reflect2"
)

func newExceptionTypeName(v any) string {
	if v == nil {
		return ""
	}
	return strings.TrimPrefix(reflect2.TypeOf(v).String(), "exception.")
}

func newSnakeCase(v string) string {
	buf := strings.Builder{}
	buf.Grow(len(v))
	for i := 0; i < len(v); i++ {
		c := rune(v[i])
		isUpper := unicode.IsUpper(c)
		if isUpper {
			c = unicode.ToLower(c)
		}
		if i > 0 && isUpper {
			buf.WriteByte('_')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
