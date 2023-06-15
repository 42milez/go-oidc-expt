package auth

import (
	"bytes"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/pkg/clock"
	"github.com/42milez/go-oidc-server/pkg/testutil/fixture"
	"testing"
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
	j, err := NewJWT(clock.RealClocker{})

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
