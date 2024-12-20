package api

import (
	"context"
	"net/url"

	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/entity"
	"github.com/42milez/go-oidc-expt/cmd/iface"
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
	VerifyConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) (bool, error)
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
	SaveAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, redirectURI, nonce, authCode string) error
}

type Authorizer interface {
	Authorize(ctx context.Context, clientID typedef.ClientID, redirectURI, state string) (*url.URL, string, error)
	RequestFingerprintSaver
}

type ConsentAcceptor interface {
	AcceptConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) error
}

//  Token
// --------------------------------------------------

type CredentialValidator interface {
	ValidateCredential(ctx context.Context, clientID typedef.ClientID, clientSecret string) error
}

type AccessTokenParser interface {
	ParseAccessToken(token string) (jwt.Token, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(ctx context.Context, token string, clientID typedef.ClientID) error
}

type UserIDExtractor interface {
	ExtractUserID(refreshToken string) (typedef.UserID, error)
}

type AuthCodeRevoker interface {
	RevokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) error
}

type TokenCacheReadWriter interface {
	iface.AuthorizationRequestFingerprintReader
	iface.RefreshTokenWriter
}

type AuthCodeGrantAcceptor interface {
	AuthCodeRevoker
	iface.AccessTokenGenerator
	iface.RefreshTokenGenerator
	iface.IDTokenGenerator
}

type RefreshTokenGrantAcceptor interface {
	RefreshTokenVerifier
	UserIDExtractor
	iface.AccessTokenGenerator
	iface.RefreshTokenGenerator
}

//  UserInfo
// --------------------------------------------------

type UserInfoReader interface {
	ReadUserInfo(ctx context.Context, accessToken jwt.Token) (*entity.UserInfo, error)
}
