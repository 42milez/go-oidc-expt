package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-server/pkg/typedef"

	"github.com/42milez/go-oidc-server/cmd/httpstore"
	"github.com/42milez/go-oidc-server/cmd/iface"
	"github.com/42milez/go-oidc-server/cmd/option"
	"github.com/42milez/go-oidc-server/cmd/repository"
	"github.com/42milez/go-oidc-server/pkg/xtime"

	"github.com/42milez/go-oidc-server/pkg/xerr"
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

func (a *AuthCodeGrant) RevokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) error {
	if err := a.validateAuthCode(ctx, code, clientID); err != nil {
		return err
	}
	if err := a.revokeAuthCode(ctx, code, clientID); err != nil {
		return err
	}
	return nil
}

func (a *AuthCodeGrant) validateAuthCode(ctx context.Context, code string, clientID typedef.ClientID) error {
	authCode, err := a.repo.ReadAuthCode(ctx, code, clientID)
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

func (a *AuthCodeGrant) revokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) error {
	_, err := a.repo.RevokeAuthCode(ctx, code, clientID)
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

func (a *AuthCodeGrant) GenerateIDToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	return generateIDToken(a.token, uid, audiences, authTime, nonce)
}

func NewRefreshTokenGrant(opt *option.Option) *RefreshTokenGrant {
	return &RefreshTokenGrant{
		cache: httpstore.NewCache(opt),
		token: opt.Token,
	}
}

type RefreshTokenGrant struct {
	cache iface.RefreshTokenReader
	token iface.TokenProcessor
}

func (r *RefreshTokenGrant) VerifyRefreshToken(ctx context.Context, token string, clientID typedef.ClientID) error {
	rt1, err := r.token.Parse(token)
	if err != nil {
		return xerr.InvalidToken
	}

	uid, err := strconv.Atoi(rt1.Subject())
	if err != nil {
		return err
	}

	rt2, err := r.cache.ReadRefreshToken(ctx, clientID, typedef.UserID(uid))
	if err != nil {
		return xerr.RefreshTokenNotFound
	}

	if !jwt.Equal(rt1, rt2) {
		return xerr.RefreshTokenNotMatched
	}

	return nil
}

func (r *RefreshTokenGrant) ExtractUserID(refreshToken string) (typedef.UserID, error) {
	t, err := r.token.Parse(refreshToken)
	if err != nil {
		return 0, err
	}
	uid, err := strconv.Atoi(t.Subject())
	if err != nil {
		return 0, err
	}
	return typedef.UserID(uid), nil
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
	idToken, err := tokenGen.GenerateIDToken(uid, audiences, authTime, nonce)
	if err != nil {
		return "", err
	}
	return idToken, nil
}
