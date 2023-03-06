package codec

import "encoding/json"

const ContentTypeJSON = "application/json"

var _ EncodeFunc = EncodeJSON

func EncodeJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}

var _ DecodeFunc = DecodeJSON

func DecodeJSON(src []byte, v any) error {
	return json.Unmarshal(src, v)
}

type JSON struct{}

var _ Codec = JSON{}

func (p JSON) Encode(v any) ([]byte, error) {
	return EncodeProtobuf(v)
}

func (p JSON) Decode(src []byte, v any) error {
	return DecodeProtobuf(src, v)
}
