package security

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
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
		t.Fatalf("%+v: %+v", xerr.FailedToInitialize, err)
	}

	wantUID := typedef.UserID(485911246986543469)
	wantUIDString := strconv.FormatUint(uint64(wantUID), 10)

	tests := map[string]struct {
		Generator func(uid typedef.UserID) (string, error)
		UserID    typedef.UserID
	}{
		"AccessToken": {
			Generator: j.GenerateAccessToken,
			UserID:    wantUID,
		},
		"RefreshToken": {
			Generator: j.GenerateRefreshToken,
			UserID:    wantUID,
		},
		"IDToken": {
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

func TestJWT_ExtractAccessToken(t *testing.T) {
	t.Parallel()

	uid := typedef.UserID(484493849344409965)
	clock := xtestutil.FixedClocker{}
	want, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(config.Issuer).
		Subject(strconv.FormatUint(uint64(uid), 10)).
		IssuedAt(clock.Now()).
		Expiration(clock.Now().Add(30 * time.Minute)).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	privateKey, err := jwk.ParseKey(rawPrivateKey, jwk.WithPEM(true))
	if err != nil {
		t.Fatal(err)
	}

	signed, err := jwt.Sign(want, jwt.WithKey(jwa.ES256, privateKey))
	if err != nil {
		t.Fatal(err)
	}

	j, err := NewJWT(clock)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "https://github.com/42milez", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signed))

	got, err := j.ExtractAccessToken(req)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want = %+v; got = %+v", want, got)
	}
}
