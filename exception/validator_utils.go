package exception

import (
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
	"oneof":           {},
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
}

func NewFromValidator(v validator.FieldError) Exception {
	tag := v.Tag()
	switch tag {
	case "required":
		return MissingParameter{
			Field: v.Field(),
		}
	case "len", "eq":
		return ParameterOutOfRange{
			Field: v.Field(),
			Len:   v.Param(),
		}
	case "gt", "gte", "max":
		return ParameterOutOfRange{
			Field: v.Field(),
			From:  v.Param(),
		}
	case "lte", "lt", "min":
		return ParameterOutOfRange{
			Field: v.Field(),
			To:    v.Param(),
		}
	default:
		// using hash set to ensure only one comparison instead N on switch case.
		if _, ok := formatSupportSet[tag]; ok {
			return InvalidParameter{
				Field:       v.Field(),
				ValidValues: v.Param(),
			}
		}
		return DomainGeneric{Parent: v.(error)}
	}
}
