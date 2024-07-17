package security

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	"github.com/42milez/go-oidc-expt/pkg/xtime"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/config"
	"github.com/42milez/go-oidc-expt/pkg/xtestutil"
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

func TestJWT_GenerateIDToken(t *testing.T) {
	t.Parallel()

	clock := &xtestutil.FixedClocker{}
	j, err := NewJWT(clock)
	if err != nil {
		t.Fatal(err)
	}
	uid := typedef.UserID(491870509865107821)

	tests := map[string]struct {
		WantUserID typedef.UserID
		WantClaims map[string]any
	}{
		"OK": {
			WantUserID: uid,
			WantClaims: map[string]any{
				jwt.IssuerKey:     config.Issuer,
				jwt.SubjectKey:    strconv.FormatUint(uint64(uid), 10),
				jwt.AudienceKey:   []string{"RZYY4jJnxBSH5vifs4bKma03wkRgee"},
				jwt.ExpirationKey: clock.Now().Add(config.IDTokenTTL),
				jwt.IssuedAtKey:   clock.Now(),
				authTimeKey:       float64(clock.Now().Unix()),
				nonceKey:          "EZeNAZyB0tXxZzUJuICiW1yqBHi3FB",
			},
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			authTimeUnix, ok := tt.WantClaims[authTimeKey].(float64)
			if !ok {
				t.Fatal("type assertion failed")
			}
			authTime := time.Unix(int64(authTimeUnix), 0)

			audiences, ok := tt.WantClaims[jwt.AudienceKey].([]string)
			if !ok {
				t.Fatal("type assertion failed")
			}

			nonce, ok := tt.WantClaims[nonceKey].(string)
			if !ok {
				t.Fatal("type assertion failed")
			}

			got, err := j.GenerateIDToken(tt.WantUserID, audiences, authTime, nonce)
			if err != nil {
				t.Fatal(err)
			}

			gotToken, err := jwt.ParseString(got, jwt.WithKey(jwa.ES256, j.publicKey), jwt.WithValidate(false))
			if err != nil {
				t.Fatal(err)
			}

			for k, claim := range tt.WantClaims {
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
		//"IDToken_OK": {
		//	Generator: j.GenerateIDToken,
		//	UserID:    uid,
		//},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			token, err := tt.Generator(tt.UserID, nil)
			if err != nil {
				t.Fatal(err)
			}

			if _, err := j.Parse(token); err != nil {
				t.Error(err)
			}
		})
	}
}
