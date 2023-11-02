package security

import (
	_ "embed"
)

//go:embed secret/key/block.key
var RawCookieBlockKey []byte

//go:embed secret/key/hash.key
var RawCookieHashKey []byte
