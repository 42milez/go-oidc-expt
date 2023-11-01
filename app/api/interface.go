package api

import (
	"context"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/typedef"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

//  Health Check
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

//  Authentication
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

type UserRegisterer interface {
	RegisterUser(ctx context.Context, name, pw string) (*entity.User, error)
}

//  Authorization
// --------------------------------------------------

type Authorizer interface {
	Authorize(ctx context.Context, clientID, redirectURI, state string) (string, string, error)
}

type ConsentAcceptor interface {
	AcceptConsent(ctx context.Context, userID typedef.UserID, clientID string) error
}

//  Token
// --------------------------------------------------

type CredentialValidator interface {
	ValidateCredential(ctx context.Context, clientID, clientSecret string) error
}

type AuthCodeValidator interface {
	ValidateAuthCode(ctx context.Context, code, clientId string) error
}

type RefreshTokenValidator interface {
	ValidateRefreshToken(token *string) error
}

type TokenRequestValidator interface {
	AuthCodeValidator
	RefreshTokenValidator
}

type AuthCodeRevoker interface {
	RevokeAuthCode(ctx context.Context, code, clientId string) error
}

type TokenSessionReader interface {
	iface.OpenIdParamSessionReader
	iface.RefreshTokenPermissionSessionReader
}

type TokenRequestAcceptor interface {
	TokenRequestValidator
	AuthCodeRevoker
	iface.AccessTokenGenerator
	iface.RefreshTokenGenerator
	iface.IdTokenGenerator
}
