package auth

import (
	_ "embed"
	"fmt"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/pkg/clock"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"time"
)

//go:embed cert/secret.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPublicKey []byte

type JWT struct {
	privateKey, publicKey jwk.Key
	clock                 clock.Clocker
}

func NewJWT(clock clock.Clocker) (*JWT, error){
	privKey, err := parseKey(rawPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	pubKey, err := parseKey(rawPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return &JWT{
		privateKey: privKey,
		publicKey:  pubKey,
		clock:      clock,
	}, nil
}

const issuer = "github.com/42milez/go-oidc-server"
const accessTokenSubject = "access_token"
const nameKey = "name"

func (p *JWT) GenerateAdminAccessToken(admin *ent.Admin) ([]byte, error) {
	j, err := jwt.
		NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(issuer).
		Subject(accessTokenSubject).
		IssuedAt(p.clock.Now().Add(30*time.Minute)).
		Claim(nameKey, admin.Name).
		Build()

	if err != nil {
		return nil, fmt.Errorf("failed to build token: %w", err)
	}

	signed, err := jwt.Sign(j, jwt.WithKey(jwa.ES256, p.privateKey))

	if err != nil {
		return nil, err
	}

	return signed, nil
}

func parseKey(key []byte) (jwk.Key, error) {
	k, err := jwk.ParseKey(key, jwk.WithPEM(true))
	if err != nil {
		return nil, fmt.Errorf("failed to parse key: %w", err)
	}
	return k, nil
}

func (p *JWT) ParseRequest(r *http.Request) (jwt.Token, error ) {
	return jwt.ParseRequest(r, jwt.WithKey(jwa.ES256, p.publicKey), jwt.WithValidate(false))
}

func (p *JWT) Validate(token jwt.Token) error {
	return jwt.Validate(token, jwt.WithClock(p.clock))
}
