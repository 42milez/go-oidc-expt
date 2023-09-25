package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/ent/ent"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

//  JWT
// --------------------------------------------------

type TokenGenerator interface {
	MakeAccessToken(name string) ([]byte, error)
}

//  SESSION
// --------------------------------------------------

type SessionCreator interface {
	Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error)
}

type SessionReader interface {
	Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error)
}

type SessionUpdater interface {
	Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error)
}

//  HEALTH CHECK
// --------------------------------------------------

type CachePingSender interface {
	PingCache(ctx context.Context) error
}

type DatabasePingSender interface {
	PingDatabase(ctx context.Context) error
}

type HealthChecker interface {
	CachePingSender
	DatabasePingSender
}

//  AUTHENTICATION
// --------------------------------------------------

type UserCreator interface {
	CreateUser(ctx context.Context, name string, pw string) (*ent.User, error)
}

type ConsentReader interface {
	ReadConsent(ctx context.Context, userID typedef.UserID, clientID string) (*ent.Consent, error)
}

type UserByNameReader interface {
	ReadUserByName(ctx context.Context, name string) (*ent.User, error)
}

type UserConsentReader interface {
	ConsentReader
	UserByNameReader
}

//  OIDC: AUTHORIZATION
// --------------------------------------------------

type AuthCodeCreator interface {
	CreateAuthCode(ctx context.Context, code string, clientID string, userID typedef.UserID) (*ent.AuthCode, error)
}

type RedirectUriByRelyingPartyIDReader interface {
	ReadRedirectUriByClientID(ctx context.Context, clientID string) ([]*ent.RedirectURI, error)
}

type Authorizer interface {
	AuthCodeCreator
	RedirectUriByRelyingPartyIDReader
}

type ConsentCreator interface {
	CreateConsent(ctx context.Context, userID typedef.UserID, clientID string) (*ent.Consent, error)
}
