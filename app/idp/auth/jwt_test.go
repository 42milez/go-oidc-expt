package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/testutil"
	"github.com/42milez/go-oidc-server/pkg/util"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN EC PRIVATE KEY-----")
	if !bytes.Contains(rawPrivateKey, want) {
		t.Errorf("invalid format: want = %s; got = %s", want, rawPrivateKey)
	}
	want = []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPublicKey, want) {
		t.Errorf("invalid format: want = %s; got = %s", want, rawPublicKey)
	}
}

func TestJWT_GenerateAccessToken(t *testing.T) {
	jwtUtil, err := NewJWTUtil(util.RealClocker{})

	if err != nil {
		t.Fatalf("%v: %v", xerr.FailedToInitialize, err)
	}

	got, err := jwtUtil.GenerateAccessToken("test_user")

	if err != nil {
		t.Fatal(err)
	}

	if len(got) == 0 {
		t.Errorf("want = ( not empty ); got = ( empty )")
	}
}

func TestJWT_ParseRequest(t *testing.T) {
	t.Parallel()

	c := testutil.FixedClocker{}

	want, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(nameKey, "test_admin").
		Build()

	if err != nil {
		t.Fatalf("%v: %v", ErrFailedToBuildToken, err)
	}

	privateKey, err := jwk.ParseKey(rawPrivateKey, jwk.WithPEM(true))

	if err != nil {
		t.Fatal(err)
	}

	signed, err := jwt.Sign(want, jwt.WithKey(jwa.ES256, privateKey))

	if err != nil {
		t.Fatal(err)
	}

	jwtUtil, err := NewJWTUtil(c)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "https://github.com/42milez", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signed))

	got, err := jwtUtil.parseRequest(req)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want = %v; got = %v", want, got)
	}
}

func TestJWT_Validate(t *testing.T) {
	t.Parallel()

	c := testutil.FixedClocker{}

	t.Run("OK", func(t *testing.T) {
		token, err := jwt.NewBuilder().
			JwtID(uuid.New().String()).
			Issuer(issuer).
			Subject(accessTokenSubject).
			IssuedAt(c.Now()).
			Expiration(c.Now().Add(30*time.Minute)).
			Claim(nameKey, "test_admin").
			Build()

		if err != nil {
			t.Fatalf("%v: %v", ErrFailedToBuildToken, err)
		}

		jwtUtil, err := NewJWTUtil(c)

		if err != nil {
			t.Fatalf("%v: %v", xerr.FailedToInitialize, err)
		}

		if err = jwtUtil.validate(token); err != nil {
			t.Error(err)
		}
	})

	t.Run("NG", func(t *testing.T) {
		t.Parallel()

		token, err := jwt.NewBuilder().
			JwtID(uuid.New().String()).
			Issuer(issuer).
			Subject(accessTokenSubject).
			IssuedAt(c.Now()).
			Expiration(c.Now().Add(30*time.Minute)).
			Claim(nameKey, "test_admin").
			Build()

		if err != nil {
			t.Fatalf("%v: %v", ErrFailedToBuildToken, err)
		}

		jwtUtil, err := NewJWTUtil(testutil.FixedTomorrowClocker{})

		if err != nil {
			t.Fatalf("%v: %v", xerr.FailedToInitialize, err)
		}

		if err = jwtUtil.validate(token); err == nil {
			t.Errorf("want = ( error ); got = nil")
		}
	})
}
