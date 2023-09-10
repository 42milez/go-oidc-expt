package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

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

func TestJWT_MakeAccessToken(t *testing.T) {
	t.Parallel()

	j, err := NewJWT(&xtime.RealClocker{})

	if err != nil {
		t.Fatalf("%+v: %+v", xerr.FailedToInitialize, err)
	}

	got, err := j.MakeAccessToken("test_user")

	if err != nil {
		t.Fatal(err)
	}

	if len(got) == 0 {
		t.Errorf("want = ( not empty ); got = ( empty )")
	}
}

func TestJWT_ExtractAccessToken(t *testing.T) {
	t.Parallel()

	clock := xtestutil.FixedClocker{}
	want, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(clock.Now()).
		Expiration(clock.Now().Add(30*time.Minute)).
		Claim(nameKey, "test_user").
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
