package auth

import (
	"bytes"
	"fmt"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/pkg/testutil/fixture"
	"github.com/42milez/go-oidc-server/pkg/util"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN EC PRIVATE KEY-----")
	if !bytes.Contains(rawPrivateKey, want) {
		t.Errorf("want = %s; got = %s", want, rawPrivateKey)
	}
	want = []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPublicKey, want) {
		t.Errorf("want = %s; got = %s", want, rawPublicKey)
	}
}

func TestJWT_GenerateAdminAccessToken(t *testing.T) {
	admin := fixture.Admin(&ent.Admin{})
	j, err := NewJWT(util.RealClocker{})

	if err != nil {
		t.Fatal(err)
	}

	got, err := j.GenerateAdminAccessToken(admin)

	if err != nil {
		t.Fatal(err)
	}

	if len(got) == 0 {
		t.Errorf("empty token")
	}
}

func TestJWT_ParseRequest(t *testing.T) {
	t.Parallel()

	c := util.FixedClocker{}

	want, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(nameKey, "test_admin").
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

	j, err := NewJWT(c)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "https://github.com/42milez", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signed))

	got, err := j.ParseRequest(req)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want = %v; got = %v", want, got)
	}
}

func TestJWT_Validate(t *testing.T) {
	t.Parallel()

	c := util.FixedClocker{}

	validToken, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(nameKey, "test_admin").
		Build()

	if err != nil {
		t.Fatal(err)
	}

	j, err := NewJWT(c)

	if err != nil {
		t.Fatal(err)
	}

	if err := j.Validate(validToken); err != nil {
		t.Error(err)
	}
}
