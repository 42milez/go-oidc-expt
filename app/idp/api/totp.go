package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/config"

	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
)

func NewTOTP(key string) *TOTP {
	return &TOTP{
		Key: key,
		now: func() int64 { return time.Now().Unix() },
	}
}

type TOTP struct {
	Key string
	now func() int64
}

func hmacSha1(k []byte, t int64) []byte {
	h := hmac.New(sha1.New, k)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t))
	h.Write(buf)
	return h.Sum(nil)
}

const otpLength = 6
const reductionModulo = 1000000 // TOTP is represented by 6-digit number.

func dt(input []byte) string {
	offset := input[19] & 0xf
	binCode := int(input[offset]&0x7f)<<24 |
		int(input[offset+1]&0xff)<<16 |
		int(input[offset+2]&0xff)<<8 |
		int(input[offset+3]&0xff)
	otp := binCode % reductionModulo
	result := strconv.Itoa(otp)

	for len(result) < otpLength {
		result = "0" + result
	}

	return result
}

const secretLength = 30
const t0 = 0
const timeStep = 30

func format(email, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", config.Issuer, email, secret, config.Issuer)
}

func SecretKey(email string) (string, error) {
	var secret string
	var e error

	if secret, e = xrandom.GenerateCryptoRandomString(secretLength); e != nil {
		return "", e
	}

	secretEnc32 := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(secret))

	return format(email, secretEnc32), nil
}

func (totp *TOTP) Verify(code string) (bool, error) {
	t := (totp.now() - t0) / timeStep

	secret, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(totp.Key)
	if err != nil {
		return false, err
	}

	hs := hmacSha1(secret, t)
	d := dt(hs)
	if d != code {
		return false, nil
	}

	return true, nil
}
