package auth

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xtime"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed cert/private.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPublicKey []byte

type JWT struct {
	privateKey, publicKey jwk.Key
	clock                 xtime.Clocker
}

func NewJWT(clock xtime.Clocker) (*JWT, error) {
	priKey, err := parseKey(rawPrivateKey)
	if err != nil {
		return nil, xerr.FailedToParsePrivateKey.Wrap(err)
	}

	pubKey, err := parseKey(rawPublicKey)
	if err != nil {
		return nil, xerr.FailedToParsePublicKey.Wrap(err)
	}

	return &JWT{
		privateKey: priKey,
		publicKey:  pubKey,
		clock:      clock,
	}, nil
}

const issuer = "github.com/42milez/go-oidc-server"
const accessTokenSubject = "access_token"
const nameKey = "name"

func (p *JWT) MakeAccessToken(name string) ([]byte, error) {
	token, err := jwt.
		NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(p.clock.Now().Add(30*time.Minute)).
		Claim(nameKey, name).
		Build()

	if err != nil {
		return nil, xerr.FailedToBuildToken.Wrap(err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, p.privateKey))

	if err != nil {
		return nil, xerr.FailedToSignToken.Wrap(err)
	}

	return signed, nil
}

func parseKey(key []byte) (jwk.Key, error) {
	return jwk.ParseKey(key, jwk.WithPEM(true))
}

func (p *JWT) parseRequest(r *http.Request) (jwt.Token, error) {
	return jwt.ParseRequest(r, jwt.WithKey(jwa.ES256, p.publicKey), jwt.WithValidate(false))
}

func (p *JWT) validate(token jwt.Token) error {
	return jwt.Validate(token, jwt.WithClock(p.clock))
}

func (p *JWT) Extract(r *http.Request) (jwt.Token, error) {
	token, err := p.parseRequest(r)

	if err != nil {
		return nil, xerr.FailedToParseRequest.Wrap(err)
	}

	if err = p.validate(token); err != nil {
		return nil, xerr.InvalidToken.Wrap(err)
	}

	return token, nil
}
