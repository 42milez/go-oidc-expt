package security

import (
	_ "embed"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed secret/keypair/private.pem
var rawPrivateKey []byte

//go:embed secret/keypair/public.pem
var rawPublicKey []byte

func NewJWT(clock iface.Clocker) (*JWT, error) {
	parseKey := func(key []byte) (jwk.Key, error) {
		return jwk.ParseKey(key, jwk.WithPEM(true))
	}

	ret := &JWT{
		clock: clock,
	}

	var err error

	if ret.privateKey, err = parseKey(rawPrivateKey); err != nil {
		return nil, err
	}

	if ret.publicKey, err = parseKey(rawPublicKey); err != nil {
		return nil, err
	}

	return ret, nil
}

type JWT struct {
	privateKey, publicKey jwk.Key
	clock                 iface.Clocker
}

func (j *JWT) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	builder := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Subject(strconv.FormatUint(uint64(uid), 10)).
		Issuer(config.Issuer).
		IssuedAt(j.clock.Now()).
		Expiration(j.clock.Now().Add(config.AccessTokenTTL))

	for k := range claims {
		switch k {
		default:
			return "", xerr.UnsupportedClaim
		}
	}

	token, err := builder.Build()
	if err != nil {
		return "", err
	}

	ret, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, j.privateKey))
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

func (j *JWT) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	builder := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Subject(strconv.FormatUint(uint64(uid), 10)).
		Issuer(config.Issuer).
		IssuedAt(j.clock.Now()).
		Expiration(j.clock.Now().Add(config.RefreshTokenTTL))

	for k := range claims {
		switch k {
		default:
			return "", xerr.UnsupportedClaim
		}
	}

	token, err := builder.Build()
	if err != nil {
		return "", err
	}

	ret, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, j.privateKey))
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

const authTimeKey = "auth_time"
const nonceKey = "nonce"

// GenerateIdToken generates ID token and returns it as string. The detail of ID token is described here:
// https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#IDToken
func (j *JWT) GenerateIdToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	builder := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(config.Issuer).
		Subject(strconv.FormatUint(uint64(uid), 10)).
		Audience(audiences).
		Expiration(j.clock.Now().Add(config.IDTokenTTL)).
		IssuedAt(j.clock.Now()).
		Claim(authTimeKey, authTime.Unix()).
		Claim(nonceKey, nonce)

	token, err := builder.Build()
	if err != nil {
		return "", err
	}

	ret, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, j.privateKey))
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

func (j *JWT) Validate(token string) error {
	t, err := jwt.ParseString(token, jwt.WithKey(jwa.ES256, j.publicKey))
	if err != nil {
		return xerr.InvalidToken
	}
	return j.validate(t)
}

func (j *JWT) validate(token jwt.Token) error {
	return jwt.Validate(token, jwt.WithClock(j.clock))
}
