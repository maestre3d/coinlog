package exception

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func NewFromValidator(v validator.FieldError) Exception {
	switch v.Tag() {
	case "required":
		return MissingParameter{
			Field: v.Field(),
		}
	case "len":
		return ParameterOutOfRange{
			Field: v.Field(),
			Len:   v.Param(),
		}
	case "gt", "gte":
		return ParameterOutOfRange{
			Field: v.Field(),
			From:  v.Param(),
		}
	case "lte", "lt":
		return ParameterOutOfRange{
			Field: v.Field(),
			To:    v.Param(),
		}
	case "url", "oneof", "iso4217":
		return InvalidParameter{
			Field:       v.Field(),
			ValidValues: v.Param(),
		}
	default:
		return DomainUnknown{Parent: errors.New(v.Error())}
	}
}
