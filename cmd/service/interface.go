package service

import (
	"context"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/entity"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

//  Health Check
// --------------------------------------------------

type CachePingSender interface {
	PingCache(ctx context.Context) error
}

type DatabasePingSender interface {
	PingDatabase(ctx context.Context) error
}

type PingSender interface {
	CachePingSender
	DatabasePingSender
}

//  Authentication
// --------------------------------------------------

type UserCreator interface {
	CreateUser(ctx context.Context, name string, pw string) (*entity.User, error)
}

type ConsentReader interface {
	ReadConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) (*entity.Consent, error)
}

type UserReader interface {
	ReadUser(ctx context.Context, name string) (*entity.User, error)
	ReadUserByID(ctx context.Context, id typedef.UserID) (*entity.User, error)
}

type UserConsentReader interface {
	ConsentReader
	UserReader
}

type CredentialReader interface {
	ReadCredential(ctx context.Context, clientID typedef.ClientID, clientSecret string) (*entity.RelyingParty, error)
}

//  Authorization
// --------------------------------------------------

type AuthCodeCreator interface {
	CreateAuthCode(ctx context.Context, code string, clientID typedef.ClientID, userID typedef.UserID) (*entity.AuthCode, error)
}

type RedirectURIsReader interface {
	ReadRedirectURIs(ctx context.Context, clientID typedef.ClientID) ([]*entity.RedirectURI, error)
}

type Authorizer interface {
	AuthCodeCreator
	RedirectURIsReader
}

type ConsentCreator interface {
	CreateConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) (*entity.Consent, error)
}

//  Token
// --------------------------------------------------

type AuthCodeReader interface {
	ReadAuthCode(ctx context.Context, code string, clientID typedef.ClientID) (*entity.AuthCode, error)
}

type AuthCodeRevoker interface {
	RevokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) (*entity.AuthCode, error)
}

type AuthCodeReadRevoker interface {
	AuthCodeReader
	AuthCodeRevoker
}

type RedirectURIReader interface {
	ReadRedirectURI(ctx context.Context, clientID typedef.ClientID) (*entity.RedirectURI, error)
}
