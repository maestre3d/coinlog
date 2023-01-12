package domain

import (
	"encoding/base64"
	"fmt"
)

// PageToken custom-type representing the page to fetch by a storage system.
//
// Depending on the underlying storage engine, the token might represent an offset, identifier of the last item read
// or similar.
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
