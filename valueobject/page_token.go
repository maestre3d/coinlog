package valueobject

import (
	"encoding/base64"
	"fmt"
)

// PageToken slice of bytes custom-type with encoding/decoding capabilities for further usage. Useful for pagination
// algorithms where page tokens are possible.
//
// Depending on the pagination algorithm, the token might represent an offset, last id read or similar.
type PageToken []byte

var _ fmt.Stringer = PageToken{}

func (p PageToken) String() string {
	return string(p)
}

func NewPageToken(src string) PageToken {
	srcCopy := []byte(src)
	dst := make(PageToken, base64.URLEncoding.EncodedLen(len(srcCopy)))
	base64.URLEncoding.Encode(dst, srcCopy)
	return dst
}

func DecodePageToken(token PageToken) string {
	dst := make([]byte, base64.URLEncoding.DecodedLen(len(string(token))))
	n, _ := base64.URLEncoding.Decode(dst, token)
	return string(dst[:n])
}
