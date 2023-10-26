package api

import (
	"context"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/app/typedef"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

//  COOKIE
// --------------------------------------------------

type CookieReader interface {
	Read(r *http.Request, name string) (string, error)
}

type CookieWriter interface {
	Write(w http.ResponseWriter, name, val string, ttl time.Duration) error
}

//  SESSION
// --------------------------------------------------

type RedirectUriSessionWriter interface {
	WriteRedirectUri(ctx context.Context, sid typedef.SessionID, uri string) error
}

type UserIdSessionWriter interface {
	WriteUserId(ctx context.Context, userId typedef.UserID) (typedef.SessionID, error)
}

type SessionWriter interface {
	RedirectUriSessionWriter
	UserIdSessionWriter
}

type SessionRestorer interface {
	Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error)
}

//  HEALTH CHECK
// --------------------------------------------------

type CacheStatusChecker interface {
	CheckCacheStatus(ctx context.Context) error
}

type DBStatusChecker interface {
	CheckDBStatus(ctx context.Context) error
}

type HealthChecker interface {
	CacheStatusChecker
	DBStatusChecker
}

//  HTTP
// --------------------------------------------------

type ContextReader interface {
	Read(ctx context.Context, key any) any
}

//  AUTHENTICATION
// --------------------------------------------------

type ConsentVerifier interface {
	VerifyConsent(ctx context.Context, userID typedef.UserID, clientID string) (bool, error)
}

type PasswordVerifier interface {
	VerifyPassword(ctx context.Context, name, pw string) (typedef.UserID, error)
}

type Authenticator interface {
	ConsentVerifier
	PasswordVerifier
}

type UserCreator interface {
	CreateUser(ctx context.Context, name, pw string) (*entity.User, error)
}

type UserReader interface {
	SelectUser(ctx context.Context) (*entity.User, error)
}

//  OIDC: AUTHORIZATION
// --------------------------------------------------

type Authorizer interface {
	Authorize(ctx context.Context, clientID, redirectURI, state string) (string, error)
}

type ConsentAcceptor interface {
	AcceptConsent(ctx context.Context, userID typedef.UserID, clientID string) error
}

// OIDC: Token
// --------------------------------------------------

type CredentialValidator interface {
	ValidateCredential(ctx context.Context, clientID, clientSecret string) error
}

type AuthCodeValidator interface {
	ValidateAuthCode(ctx context.Context, code, clientId string) error
}

type RedirectUriValidator interface {
	ValidateRedirectUri(ctx context.Context, uri, clientId string) error
}

type RefreshTokenValidator interface {
	ValidateRefreshToken(token *string) error
}

type TokenRequestValidator interface {
	AuthCodeValidator
	RedirectUriValidator
	RefreshTokenValidator
}

type AuthCodeRevoker interface {
	RevokeAuthCode(ctx context.Context, code, clientId string) error
}

type AccessTokenGenerator interface {
	GenerateAccessToken(uid typedef.UserID) (string, error)
}

type RefreshTokenGenerator interface {
	GenerateRefreshToken(uid typedef.UserID) (string, error)
}

type IdTokenGenerator interface {
	GenerateIdToken(uid typedef.UserID) (string, error)
}

type TokenRequestAcceptor interface {
	TokenRequestValidator
	AuthCodeRevoker
	AccessTokenGenerator
	RefreshTokenGenerator
	IdTokenGenerator
}
