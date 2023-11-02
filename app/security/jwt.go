package security

import (
	_ "embed"
	"net/http"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/config"

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

func (j *JWT) GenerateAccessToken() (string, error) {
	return xrandom.GenerateCryptoRandomString(config.AccessTokenLength)
}

func (j *JWT) GenerateRefreshToken() (string, error) {
	return xrandom.GenerateCryptoRandomString(config.RefreshTokenLength)
}

func (j *JWT) GenerateIdToken(uid typedef.UserID) (string, error) {
	token, err := jwt.NewBuilder().JwtID(uuid.New().String()).Issuer(config.Issuer).
		Subject(strconv.FormatUint(uint64(uid), 10)).IssuedAt(j.clock.Now()).Expiration(j.clock.Now().
		Add(30 * time.Minute)).Build()
	if err != nil {
		return "", err
	}
	ret, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, j.privateKey))
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func (j *JWT) ExtractAccessToken(r *http.Request) (jwt.Token, error) {
	ret, err := j.parseRequest(r)
	if err != nil {
		return nil, err
	}
	if err = j.validate(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (j *JWT) Validate(token *string) error {
	t, err := jwt.ParseString(*token, jwt.WithKey(jwa.ES256, j.publicKey))
	if err != nil {
		return xerr.InvalidToken
	}
	return j.validate(t)
}

func (j *JWT) parseRequest(r *http.Request) (jwt.Token, error) {
	return jwt.ParseRequest(r, jwt.WithKey(jwa.ES256, j.publicKey), jwt.WithValidate(false))
}

func (j *JWT) validate(token jwt.Token) error {
	return jwt.Validate(token, jwt.WithClock(j.clock))
}
