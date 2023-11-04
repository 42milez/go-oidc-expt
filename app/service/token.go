package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/option"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewToken(opt *option.Option) *Token {
	return &Token{
		acRepo:  repository.NewAuthCode(opt.DB),
		ruRepo:  repository.NewRedirectUri(opt.DB),
		cache:   httpstore.NewCache(opt),
		clock:   &xtime.RealClocker{},
		context: &httpstore.Context{},
		token:   opt.Token,
		v:       opt.V,
	}
}

type Token struct {
	acRepo  AuthCodeReadRevoker
	ruRepo  RedirectUriReader
	cache   iface.RefreshTokenPermissionReader
	clock   iface.Clocker
	context iface.ContextReader
	token   iface.TokenGenerateValidator
	v       iface.StructValidator
}

func (t *Token) ReadRefreshTokenPermission(ctx context.Context, token, clientId string) (*typedef.RefreshTokenPermission, error) {
	if err := t.token.Validate(token); err != nil {
		return nil, xerr.InvalidToken
	}

	perm, err := t.cache.ReadRefreshTokenPermission(ctx, token)
	if err != nil {
		return nil, xerr.RefreshTokenPermissionNotFound
	}

	if perm.ClientId != clientId {
		return nil, xerr.ClientIdNotMatched
	}

	return perm, nil
}

func (t *Token) ValidateAuthCode(ctx context.Context, code, clientId string) error {
	authCode, err := t.acRepo.ReadAuthCode(ctx, code, clientId)
	if err != nil {
		return err
	}

	if !authCode.ExpireAt().After(t.clock.Now()) {
		return xerr.AuthCodeExpired
	}

	if authCode.UsedAt() != nil {
		return xerr.AuthCodeUsed
	}

	return nil
}

func (t *Token) RevokeAuthCode(ctx context.Context, code, clientId string) error {
	_, err := t.acRepo.RevokeAuthCode(ctx, code, clientId)
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) GenerateAccessToken(uid typedef.UserID) (string, error) {
	accessToken, err := t.token.GenerateAccessToken(uid)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (t *Token) GenerateRefreshToken(uid typedef.UserID) (string, error) {
	refreshToken, err := t.token.GenerateRefreshToken(uid)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (t *Token) GenerateIdToken(uid typedef.UserID) (string, error) {
	idToken, err := t.token.GenerateIdToken(uid)
	if err != nil {
		return "", err
	}
	return idToken, nil
}
