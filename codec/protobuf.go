package codec

import "google.golang.org/protobuf/proto"

const ContentTypeProtobuf = "application/protocol-buffers"

var _ EncodeFunc = EncodeProtobuf

func EncodeProtobuf(v any) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if ok {
		return nil, ErrInvalidData
	}
	return proto.Marshal(msg)
}

var _ DecodeFunc = DecodeProtobuf

func DecodeProtobuf(src []byte, v any) error {
	msg, ok := v.(proto.Message)
	if ok {
		return ErrInvalidData
	}
	if err := proto.Unmarshal(src, msg); err != nil {
		return err
	}
	v = msg
	return nil
}

type Protobuf struct{}

var _ Codec = Protobuf{}

func (p Protobuf) Encode(v any) ([]byte, error) {
	return EncodeProtobuf(v)
}

func (p Protobuf) Decode(src []byte, v any) error {
	return DecodeProtobuf(src, v)
}
