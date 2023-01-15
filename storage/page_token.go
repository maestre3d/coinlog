package storage

import (
	"encoding/base64"
	"fmt"
)

// PageToken custom-type representing the page to fetch by a storage system.
//
// This type obfuscates raw tokens through a codec to be later served to system users securely.
//
// Depending on the underlying storage engine, the token might represent an offset, cursor (identifier of an item)
// or similar.
type PageToken []byte

var _ fmt.Stringer = PageToken{}

// NewPageToken allocates a new instance of PageToken based on src.
func NewPageToken(src string) PageToken {
	srcCopy := []byte(src)
	dst := make(PageToken, base64.URLEncoding.EncodedLen(len(srcCopy)))
	base64.URLEncoding.Encode(dst, srcCopy)
	return dst
}

// DecodePageToken retrieves raw token from a PageToken.
func DecodePageToken(token PageToken) string {
	dst := make([]byte, base64.URLEncoding.DecodedLen(len(string(token))))
	n, _ := base64.URLEncoding.Decode(dst, token)
	return string(dst[:n])
}

func (p PageToken) String() string {
	return string(p)
}
