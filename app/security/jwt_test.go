package security

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
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

	clock := &xtestutil.FixedClocker{}
	j, err := NewJWT(clock)
	if err != nil {
		t.Fatal(err)
	}
	uid := typedef.UserID(485911246986543469)

	tests := map[string]struct {
		Generator           func(uid typedef.UserID, claims map[string]any) (string, error)
		UserID              typedef.UserID
		WantCommonClaims    map[string]any
		WantDedicatedClaims map[string]any
	}{
		"AccessToken_OK": {
			Generator: j.GenerateAccessToken,
			UserID:    uid,
			WantCommonClaims: map[string]any{
				jwt.IssuerKey:     config.Issuer,
				jwt.SubjectKey:    strconv.FormatUint(uint64(uid), 10),
				jwt.IssuedAtKey:   clock.Now(),
				jwt.ExpirationKey: clock.Now().Add(config.AccessTokenTTL),
			},
			WantDedicatedClaims: nil,
		},
		"RefreshToken_OK": {
			Generator: j.GenerateRefreshToken,
			UserID:    uid,
			WantCommonClaims: map[string]any{
				jwt.IssuerKey:     config.Issuer,
				jwt.SubjectKey:    strconv.FormatUint(uint64(uid), 10),
				jwt.IssuedAtKey:   clock.Now(),
				jwt.ExpirationKey: clock.Now().Add(config.RefreshTokenTTL),
			},
			WantDedicatedClaims: nil,
		},
		"IDToken_OK": {
			Generator: j.GenerateIdToken,
			UserID:    uid,
			WantCommonClaims: map[string]any{
				jwt.IssuerKey:     config.Issuer,
				jwt.SubjectKey:    strconv.FormatUint(uint64(uid), 10),
				jwt.IssuedAtKey:   clock.Now(),
				jwt.ExpirationKey: clock.Now().Add(config.IDTokenTTL),
			},
			// https://openid.net/specs/openid-connect-core-1_0.html#IDToken
			WantDedicatedClaims: map[string]any{
				jwt.AudienceKey: []string{
					"RZYY4jJnxBSH5vifs4bKma03wkRgee",
				},
				nonceKey: "EZeNAZyB0tXxZzUJuICiW1yqBHi3FB",
			},
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			got, err := tt.Generator(tt.UserID, tt.WantDedicatedClaims)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) == 0 {
				t.Fatal("want = ( not empty ); got = ( empty )")
			}

			gotToken, err := jwt.ParseString(got, jwt.WithKey(jwa.ES256, j.publicKey), jwt.WithValidate(false))
			if err != nil {
				t.Fatal(err)
			}

			for k, claim := range tt.WantCommonClaims {
				gotClaim, ok := gotToken.Get(k)
				if !ok {
					t.Fatalf("claim not included: %s", k)
				}
				xtestutil.CompareType(t, claim, gotClaim)
				xtestutil.CompareValue(t, claim, gotClaim)
			}

			for k, claim := range tt.WantDedicatedClaims {
				gotClaim, ok := gotToken.Get(k)
				if !ok {
					t.Fatalf("claim not included: %s", k)
				}
				xtestutil.CompareType(t, claim, gotClaim)
				xtestutil.CompareValue(t, claim, gotClaim)
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
		Generator func(uid typedef.UserID, claims map[string]any) (string, error)
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

			token, err := tt.Generator(tt.UserID, nil)
			if err != nil {
				t.Fatal(err)
			}

			if err := j.Validate(token); err != nil {
				t.Error(err)
			}
		})
	}
}
