package iface

import (
	"context"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/typedef"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=iface.go -destination=iface_mock.go -package=$GOPACKAGE

//  Context
// --------------------------------------------------

type ContextReader interface {
	Read(ctx context.Context, key any) any
}

//  Cookie
// --------------------------------------------------

type CookieReader interface {
	Read(r *http.Request, name string) (string, error)
}

type CookieWriter interface {
	Write(w http.ResponseWriter, name, val string, ttl time.Duration) error
}

//  JWT
// --------------------------------------------------

type AccessTokenGenerator interface {
	GenerateAccessToken(uid typedef.UserID) (string, error)
}

type RefreshTokenGenerator interface {
	GenerateRefreshToken(uid typedef.UserID) (string, error)
}

type IdTokenGenerator interface {
	GenerateIdToken(uid typedef.UserID) (string, error)
}

type TokenGenerator interface {
	AccessTokenGenerator
	RefreshTokenGenerator
	IdTokenGenerator
}

type TokenValidator interface {
	Validate(name *string) error
}

type TokenGenerateValidator interface {
	TokenGenerator
	TokenValidator
}

//  Session
// --------------------------------------------------

type RedirectUriSessionReader interface {
	ReadRedirectUri(ctx context.Context, clientId, authCode string) (string, error)
}

type RedirectUriSessionWriter interface {
	WriteRedirectUriAssociation(ctx context.Context, uri, clientId, authCode string) error
}

type RefreshTokenOwnerSessionWriter interface {
	WriteRefreshTokenOwner(ctx context.Context, token, clientId string) error
}

type UserIdSessionWriter interface {
	WriteUserId(ctx context.Context, userId typedef.UserID) (typedef.SessionID, error)
}

//  Validator
// --------------------------------------------------

type StructValidator interface {
	Struct(s any) error
}
