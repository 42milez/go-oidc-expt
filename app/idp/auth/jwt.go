package auth

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed cert/jwt/private.pem
var rawPrivateKey []byte

//go:embed cert/jwt/public.pem
var rawPublicKey []byte

const (
	errFailedToBuildToken      xerr.Err = "failed to build token"
	errFailedToParsePrivateKey xerr.Err = "failed to parse private key"
	errFailedToParsePublicKey  xerr.Err = "failed to parse public key"
	errFailedToParseRequest    xerr.Err = "failed to parse request"
	errFailedToSignToken       xerr.Err = "failed to sign token"
	errInvalidToken            xerr.Err = "invalid token"
)

type JWTUtil struct {
	privateKey, publicKey jwk.Key
	clock                 xutil.Clocker
}

func NewJWTUtil(clock xutil.Clocker) (*JWTUtil, error) {
	privKey, err := parseKey(rawPrivateKey)
	if err != nil {
		return nil, xerr.Wrap(errFailedToParsePrivateKey, err)
	}

	pubKey, err := parseKey(rawPublicKey)
	if err != nil {
		return nil, xerr.Wrap(errFailedToParsePublicKey, err)
	}

	return &JWTUtil{
		privateKey: privKey,
		publicKey:  pubKey,
		clock:      clock,
	}, nil
}

const issuer = "github.com/42milez/go-oidc-server"
const accessTokenSubject = "access_token"
const nameKey = "name"

func (p *JWTUtil) GenerateAccessToken(name string) ([]byte, error) {
	token, err := jwt.
		NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(p.clock.Now().Add(30*time.Minute)).
		Claim(nameKey, name).
		Build()

	if err != nil {
		return nil, xerr.Wrap(errFailedToBuildToken, err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, p.privateKey))

	if err != nil {
		return nil, xerr.Wrap(errFailedToSignToken, err)
	}

	return signed, nil
}

func parseKey(key []byte) (jwk.Key, error) {
	return jwk.ParseKey(key, jwk.WithPEM(true))
}

func (p *JWTUtil) parseRequest(r *http.Request) (jwt.Token, error) {
	return jwt.ParseRequest(r, jwt.WithKey(jwa.ES256, p.publicKey), jwt.WithValidate(false))
}

func (p *JWTUtil) validate(token jwt.Token) error {
	return jwt.Validate(token, jwt.WithClock(p.clock))
}

func (p *JWTUtil) ExtractToken(r *http.Request) (jwt.Token, error) {
	token, err := p.parseRequest(r)

	if err != nil {
		return nil, xerr.Wrap(errFailedToParseRequest, err)
	}

	if err = p.validate(token); err != nil {
		return nil, xerr.Wrap(errInvalidToken, err)
	}

	return token, nil
}
