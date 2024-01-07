package api

import (
	"context"
	"net/url"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/iface"
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

type RequestFingerprintSaver interface {
	SaveAuthorizationRequestFingerprint(ctx context.Context, clientID, redirectURI, nonce, authCode string) error
}

type Authorizer interface {
	Authorize(ctx context.Context, clientID, redirectURI, state string) (*url.URL, string, error)
	RequestFingerprintSaver
}

type ConsentAcceptor interface {
	AcceptConsent(ctx context.Context, userID typedef.UserID, clientID string) error
}

//  Token
// --------------------------------------------------

type CredentialValidator interface {
	ValidateCredential(ctx context.Context, clientID, clientSecret string) error
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(ctx context.Context, token string, clientId string) error
}

type UserIDExtractor interface {
	ExtractUserID(refreshToken string) (typedef.UserID, error)
}

type AuthCodeRevoker interface {
	RevokeAuthCode(ctx context.Context, code, clientId string) error
}

type TokenCacheReadWriter interface {
	iface.AuthorizationRequestFingerprintReader
	iface.RefreshTokenWriter
}

type AuthCodeGrantAcceptor interface {
	AuthCodeRevoker
	iface.AccessTokenGenerator
	iface.RefreshTokenGenerator
	iface.IdTokenGenerator
}

type RefreshTokenGrantAcceptor interface {
	RefreshTokenVerifier
	UserIDExtractor
	iface.AccessTokenGenerator
	iface.RefreshTokenGenerator
}
