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
	}
}

type Token struct {
	acRepo  AuthCodeReadRevoker
	ruRepo  RedirectUriReader
	cache   iface.RefreshTokenOwnerReader
	clock   iface.Clocker
	context iface.ContextReader
	token   iface.TokenGenerateValidator
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

func (t *Token) ValidateRefreshToken(ctx context.Context, token *string, clientId string) error {
	ownerId, err := t.cache.ReadRefreshTokenOwner(ctx, *token)
	if err != nil {
		return err
	}
	if ownerId != clientId {
		return xerr.RefreshTokenOwnerIdNotMatched
	}
	return nil
}

func (t *Token) GenerateAccessToken() (string, error) {
	accessToken, err := t.token.GenerateAccessToken()
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (t *Token) GenerateRefreshToken() (string, error) {
	refreshToken, err := t.token.GenerateRefreshToken()
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
