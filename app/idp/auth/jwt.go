package auth

import (
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/pkg/util"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed cert/secret.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPublicKey []byte

const (
	ErrFailedToBuildToken      JWTErr = "failed to build token"
	ErrFailedToParseToken      JWTErr = "failed to parse token"
	ErrFailedToParsePrivateKey JWTErr = "failed to parse private key"
	ErrFailedToParsePublicKey  JWTErr = "failed to parse public key"
	ErrFailedToParseRequest    JWTErr = "failed to parse request"
	ErrFailedToSignToken       JWTErr = "failed to sign token"
	ErrInvalidToken            JWTErr = "invalid token"
)

type JWTErr string

func (v JWTErr) Error() string {
	return string(v)
}

type JWTUtil struct {
	privateKey, publicKey jwk.Key
	clock                 util.Clocker
}

func NewJWTUtil(clock util.Clocker) (*JWTUtil, error) {
	privKey, err := parseKey(rawPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParsePrivateKey, err)
	}

	pubKey, err := parseKey(rawPublicKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParsePublicKey, err)
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
		return nil, fmt.Errorf("%w: %w", ErrFailedToBuildToken, err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, p.privateKey))

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToSignToken, err)
	}

	return signed, nil
}

func parseKey(key []byte) (jwk.Key, error) {
	return jwk.ParseKey(key, jwk.WithPEM(true))
}

func (p *JWTUtil) Parse(signed []byte) (jwt.Token, error) {
	ret, err := jwt.Parse(signed, jwt.WithKey(jwa.ES256, p.publicKey))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParseToken, err)
	}
	return ret, nil
}

func (p *JWTUtil) ParseRequest(r *http.Request) (jwt.Token, error) {
	ret, err := jwt.ParseRequest(r, jwt.WithKey(jwa.ES256, p.publicKey), jwt.WithValidate(false))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParseRequest, err)
	}
	return ret, nil
}

func (p *JWTUtil) Validate(token jwt.Token) error {
	if err := jwt.Validate(token, jwt.WithClock(p.clock)); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidToken, err)
	}
	return nil
}

func (p *JWTUtil) ExtractToken(r *http.Request) (jwt.Token, error) {
	token, err := p.ParseRequest(r)

	if err != nil {
		return nil, err
	}

	if err = p.Validate(token); err != nil {
		return nil, err
	}

	return token, nil
}
