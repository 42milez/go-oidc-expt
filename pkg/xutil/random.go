package xutil

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // Number of the set of letters fitting in 63 bits
)

const int64Max = int64(^uint64(0) >> 1)

// MakeCryptoRandomString generates cryptographically secure random string.
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986
func MakeCryptoRandomString(n int) (string, error) {
	var val *big.Int
	var err error

	if val, err = randomInt(int64Max); err != nil {
		return "", err
	}

	ret := make([]byte, n)

	for i, cache, remain := n-1, val.Int64(), letterIdxMax; i >= 0; {
		if remain == 0 {
			if val, err = randomInt(int64Max); err != nil {
				return "", err
			}
			cache = val.Int64()
			remain = letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			ret[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&ret)), nil
}

// MakeCryptoRandomStringNoCache generates cryptographically secure random string without caching.
func MakeCryptoRandomStringNoCache(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		val, err := randomInt(int64(len(letterBytes)) - 1)
		if err != nil {
			return "", err
		}
		ret[i] = letterBytes[val.Int64()]
	}
	return *(*string)(unsafe.Pointer(&ret)), nil
}

func randomInt(n int64) (*big.Int, error) {
	return crand.Int(crand.Reader, big.NewInt(n))
}

var src = mrand.NewSource(time.Now().UnixNano())

// MakeMathRandomString generates random string.
func MakeMathRandomString(n int) string {
	ret := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			ret[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&ret))
}
