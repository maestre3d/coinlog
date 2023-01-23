package exception

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var formatSupportSet = map[string]struct{}{
	"url":             {},
	"url_encoded":     {},
	"uri":             {},
	"alpha":           {},
	"alphanum":        {},
	"alphanumunicode": {},
	"alphaunicode":    {},
	"ascii":           {},
	"iso4217":         {},
	"email":           {},
	"datetime":        {},
	"timezone":        {},
	"uuid":            {},
	"ipv4":            {},
	"ip":              {},
	"cidr":            {},
	"fqdn":            {},
	"mac":             {},
	"hostname":        {},
	"md5":             {},
	"sha256":          {},
	"uppercase":       {},
	"lowercase":       {},
}

func NewFromValidator(v validator.FieldError) Exception {
	tag := v.Tag()
	field := newSnakeCase(v.Field())
	switch tag {
	case "required":
		return MissingParameter{
			Field: field,
		}
	case "len", "eq":
		return ParameterOutOfRange{
			Field: field,
			Len:   v.Param(),
		}
	case "gt", "gte", "max":
		return ParameterOutOfRange{
			Field: field,
			From:  v.Param(),
		}
	case "lte", "lt", "min":
		return ParameterOutOfRange{
			Field: field,
			To:    v.Param(),
		}
	case "oneof":
		return InvalidParameter{
			Field:       field,
			ValidValues: strings.ReplaceAll(v.Param(), " ", ","),
		}
	default:
		// using hash set to ensure only one comparison instead N on switch case.
		if _, ok := formatSupportSet[tag]; ok {
			return InvalidParameter{
				Field:       field,
				ValidValues: tag,
			}
		}
		return DomainGeneric{
			Code:   tag,
			Parent: v.(error),
		}
	}
}
