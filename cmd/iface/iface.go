package iface

import (
	"context"
	"net/http"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-server/pkg/typedef"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=iface.go -destination=iface_mock.go -package=$GOPACKAGE

//  Context
// --------------------------------------------------

type ContextReader interface {
	Read(ctx context.Context, key any) any
}

//  Time
// --------------------------------------------------

type Clocker interface {
	Now() time.Time
}

//  Cookie
// --------------------------------------------------

type CookieReader interface {
	Read(r *http.Request, name string) (string, error)
}

type CookieWriter interface {
	Write(w http.ResponseWriter, name, val string, ttl time.Duration) error
}

type CookieReadWriter interface {
	CookieReader
	CookieWriter
}

//  JWT
// --------------------------------------------------

type AccessTokenGenerator interface {
	GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error)
}

type RefreshTokenGenerator interface {
	GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error)
}

type IDTokenGenerator interface {
	GenerateIDToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error)
}

type TokenGenerator interface {
	AccessTokenGenerator
	RefreshTokenGenerator
	IDTokenGenerator
}

type TokenParser interface {
	Parse(token string) (jwt.Token, error)
}

type TokenProcessor interface {
	TokenGenerator
	TokenParser
}

//  Cache
// --------------------------------------------------

type AuthorizationRequestFingerprintReader interface {
	ReadAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string) (*typedef.AuthorizationRequestFingerprint, error)
}

type AuthorizationRequestFingerprintWriter interface {
	WriteAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string, param *typedef.AuthorizationRequestFingerprint) error
}

type RefreshTokenReader interface {
	ReadRefreshToken(ctx context.Context, clientID typedef.ClientID, userID typedef.UserID) (jwt.Token, error)
}

type RefreshTokenWriter interface {
	WriteRefreshToken(ctx context.Context, token string, clientID typedef.ClientID, userID typedef.UserID) error
}

type SessionCreator interface {
	CreateSession(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error)
}

//  Validator
// --------------------------------------------------

type StructValidator interface {
	Struct(s any) error
}

//  ID Generator
// --------------------------------------------------

type IDGenerator interface {
	NextID() (uint64, error)
}
