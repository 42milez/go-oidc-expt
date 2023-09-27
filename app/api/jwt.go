package api

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const issuer = "github.com/42milez/go-oidc-server"
const accessTokenSubject = "access_token"
const nameKey = "name"

//go:embed secret/keypair/private.pem
var rawPrivateKey []byte

//go:embed secret/keypair/public.pem
var rawPublicKey []byte

func NewJWT(clock xtime.Clocker) (*JWT, error) {
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
	clock                 xtime.Clocker
}

func (j *JWT) ExtractAccessToken(r *http.Request) (jwt.Token, error) {
	var ret jwt.Token
	var err error

	if ret, err = j.parseRequest(r); err != nil {
		return nil, err
	}

	if err = j.validate(ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (j *JWT) MakeAccessToken(name string) ([]byte, error) {
	var token jwt.Token
	var err error

	if token, err = jwt.NewBuilder().JwtID(uuid.New().String()).Issuer(issuer).Subject(accessTokenSubject).
		IssuedAt(j.clock.Now().Add(30*time.Minute)).Claim(nameKey, name).Build(); err != nil {
		return nil, err
	}

	var ret []byte

	if ret, err = jwt.Sign(token, jwt.WithKey(jwa.ES256, j.privateKey)); err != nil {
		return nil, err
	}

	return ret, nil
}

func (j *JWT) parseRequest(r *http.Request) (jwt.Token, error) {
	return jwt.ParseRequest(r, jwt.WithKey(jwa.ES256, j.publicKey), jwt.WithValidate(false))
}

func (j *JWT) validate(token jwt.Token) error {
	return jwt.Validate(token, jwt.WithClock(j.clock))
}
