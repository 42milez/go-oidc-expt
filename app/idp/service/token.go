package service

import (
	"context"
	"errors"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewAuthCodeGrant(opt *option.Option) *AuthCodeGrant {
	return &AuthCodeGrant{
		repo:  repository.NewAuthCode(opt.DB),
		clock: &xtime.RealClocker{},
		token: opt.Token,
	}
}

type AuthCodeGrant struct {
	repo  AuthCodeReadRevoker
	clock iface.Clocker
	token iface.TokenGenerator
}

func (a *AuthCodeGrant) RevokeAuthCode(ctx context.Context, code, clientId string) error {
	if err := a.validateAuthCode(ctx, code, clientId); err != nil {
		return err
	}
	if err := a.revokeAuthCode(ctx, code, clientId); err != nil {
		return err
	}
	return nil
}

func (a *AuthCodeGrant) validateAuthCode(ctx context.Context, code, clientId string) error {
	authCode, err := a.repo.ReadAuthCode(ctx, code, clientId)
	if err != nil {
		if errors.Is(err, xerr.RecordNotFound) {
			return xerr.AuthCodeNotFound
		} else {
			return err
		}
	}

	if !authCode.ExpireAt().After(a.clock.Now()) {
		return xerr.AuthCodeExpired
	}

	if authCode.UsedAt() != nil {
		return xerr.AuthCodeUsed
	}

	return nil
}

func (a *AuthCodeGrant) revokeAuthCode(ctx context.Context, code, clientId string) error {
	_, err := a.repo.RevokeAuthCode(ctx, code, clientId)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthCodeGrant) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	return generateAccessToken(a.token, uid, claims)
}

func (a *AuthCodeGrant) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	return generateRefreshToken(a.token, uid, claims)
}

func (a *AuthCodeGrant) GenerateIdToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	return generateIDToken(a.token, uid, audiences, authTime, nonce)
}

func NewRefreshTokenGrant(opt *option.Option) *RefreshTokenGrant {
	return &RefreshTokenGrant{
		cache: httpstore.NewCache(opt),
		token: opt.Token,
	}
}

type RefreshTokenGrant struct {
	cache iface.RefreshTokenPermissionReader
	token iface.TokenGenerateValidator
}

func (r *RefreshTokenGrant) ReadRefreshTokenPermission(ctx context.Context, token, clientId string) (*typedef.RefreshTokenPermission, error) {
	if err := r.token.Validate(token); err != nil {
		return nil, xerr.InvalidToken
	}

	perm, err := r.cache.ReadRefreshTokenPermission(ctx, token)
	if err != nil {
		return nil, xerr.RefreshTokenPermissionNotFound
	}

	if perm.ClientId != clientId {
		return nil, xerr.ClientIdNotMatched
	}

	return perm, nil
}

func (r *RefreshTokenGrant) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	return generateAccessToken(r.token, uid, claims)
}

func (r *RefreshTokenGrant) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	return generateRefreshToken(r.token, uid, claims)
}

func generateAccessToken(tokenGen iface.TokenGenerator, uid typedef.UserID, claims map[string]any) (string, error) {
	accessToken, err := tokenGen.GenerateAccessToken(uid, claims)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func generateRefreshToken(tokenGen iface.TokenGenerator, uid typedef.UserID, claims map[string]any) (string, error) {
	refreshToken, err := tokenGen.GenerateRefreshToken(uid, claims)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func generateIDToken(tokenGen iface.TokenGenerator, uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	idToken, err := tokenGen.GenerateIdToken(uid, audiences, authTime, nonce)
	if err != nil {
		return "", err
	}
	return idToken, nil
}
