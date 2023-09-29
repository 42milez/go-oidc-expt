package api

import (
	"context"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/app/typedef"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

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

type SessionCreator interface {
	Create(ctx context.Context, sess *entity.Session) (string, error)
}

type SessionRestorer interface {
	Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error)
}

type SessionUpdater interface {
	Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) error
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
	CreateUser(ctx context.Context, name, pw string) (*ent.User, error)
}

type UserReader interface {
	SelectUser(ctx context.Context) (*ent.User, error)
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
