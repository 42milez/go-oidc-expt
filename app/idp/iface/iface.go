package iface

import (
	"context"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"
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

type IdTokenGenerator interface {
	GenerateIdToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error)
}

type TokenGenerator interface {
	AccessTokenGenerator
	RefreshTokenGenerator
	IdTokenGenerator
}

type TokenValidator interface {
	Validate(name string) error
}

type TokenGenerateValidator interface {
	TokenGenerator
	TokenValidator
}

//  Cache
// --------------------------------------------------

type OpenIdParamReader interface {
	ReadOpenIdParam(ctx context.Context, clientId, authCode string) (*typedef.OIDCParam, error)
}

type OpenIdParamWriter interface {
	WriteOpenIdParam(ctx context.Context, param *typedef.OIDCParam, clientId, authCode string) error
}

type RefreshTokenPermissionReader interface {
	ReadRefreshTokenPermission(ctx context.Context, token string) (*typedef.RefreshTokenPermission, error)
}

type RefreshTokenPermissionWriter interface {
	WriteRefreshTokenPermission(ctx context.Context, token, clientId string, userId typedef.UserID) error
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

type IdGenerator interface {
	NextID() (uint64, error)
}
