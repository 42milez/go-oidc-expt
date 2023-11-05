package security

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func TestJWT_Embed(t *testing.T) {
	t.Parallel()

	want := []byte("-----BEGIN EC PRIVATE KEY-----")

	if !bytes.Contains(rawPrivateKey, want) {
		t.Errorf("invalid format: want = %s; got = %s", want, rawPrivateKey)
	}

	want = []byte("-----BEGIN PUBLIC KEY-----")

	if !bytes.Contains(rawPublicKey, want) {
		t.Errorf("invalid format: want = %s; got = %s", want, rawPublicKey)
	}
}

func TestJWT_GenerateToken(t *testing.T) {
	t.Parallel()

	j, err := NewJWT(&xtime.RealClocker{})
	if err != nil {
		t.Fatal(err)
	}

	wantUID := typedef.UserID(485911246986543469)
	wantUIDString := strconv.FormatUint(uint64(wantUID), 10)

	tests := map[string]struct {
		Generator func(uid typedef.UserID) (string, error)
		UserID    typedef.UserID
	}{
		"AccessToken_OK": {
			Generator: j.GenerateAccessToken,
			UserID:    wantUID,
		},
		"RefreshToken_OK": {
			Generator: j.GenerateRefreshToken,
			UserID:    wantUID,
		},
		"IDToken_OK": {
			Generator: j.GenerateIdToken,
			UserID:    wantUID,
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			got, err := tt.Generator(wantUID)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) == 0 {
				t.Fatal("want = ( not empty ); got = ( empty )")
			}

			gotJWT, err := jwt.ParseString(got, jwt.WithKey(jwa.ES256, j.publicKey))
			if err != nil {
				t.Fatal(err)
			}

			gotUID := gotJWT.Subject()
			if gotUID != wantUIDString {
				t.Fatalf("want = %d; got = %s", wantUID, gotUID)
			}
		})
	}
}

func TestJWT_Validate(t *testing.T) {
	t.Parallel()

	j, err := NewJWT(&xtime.RealClocker{})
	if err != nil {
		t.Fatal(err)
	}

	uid := typedef.UserID(485911246986543469)

	tests := map[string]struct {
		Generator func(uid typedef.UserID) (string, error)
		UserID    typedef.UserID
	}{
		"AccessToken_OK": {
			Generator: j.GenerateAccessToken,
			UserID:    uid,
		},
		"RefreshToken_OK": {
			Generator: j.GenerateRefreshToken,
			UserID:    uid,
		},
		"IDToken_OK": {
			Generator: j.GenerateIdToken,
			UserID:    uid,
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			token, err := tt.Generator(tt.UserID)
			if err != nil {
				t.Fatal(err)
			}

			if err := j.Validate(token); err != nil {
				t.Error(err)
			}
		})
	}
}
