package codec

import "errors"

var (
	ErrInvalidData = errors.New("codec: invalid data")
)

type (
	Codec interface {
		Encode(v any) ([]byte, error)
		Decode(src []byte, v any) error
	}
	EncodeFunc func(v any) ([]byte, error)
	DecodeFunc func(src []byte, v any) error
)
