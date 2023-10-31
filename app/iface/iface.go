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

type AuthParamSessionReader interface {
	ReadAuthParam(ctx context.Context, clientId, authCode string) (*typedef.AuthParam, error)
}

type AuthParamSessionWriter interface {
	WriteAuthParam(ctx context.Context, param *typedef.AuthParam, clientId, authCode string) error
}

type RefreshTokenPermissionSessionReader interface {
	ReadRefreshTokenPermission(ctx context.Context, token string) (*typedef.AuthParam, error)
}

type RefreshTokenPermissionSessionWriter interface {
	WriteRefreshTokenPermission(ctx context.Context, token, clientId string, uid typedef.UserID) error
}

type UserInfoSessionWriter interface {
	WriteUserInfo(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error)
}

//  Validator
// --------------------------------------------------

type StructValidator interface {
	Struct(s any) error
}
