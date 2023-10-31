package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewToken(db *datastore.Database, cache *datastore.Cache, clock xtime.Clocker, token iface.TokenGenerateValidator) *Token {
	return &Token{
		acRepo: repository.NewAuthCode(db),
		ruRepo: repository.NewRedirectUri(db),
		clock:  clock,
		ctx:    &httpstore.Context{},
		token:  token,
	}
}

type Token struct {
	acRepo AuthCodeReadRevoker
	ruRepo RedirectUriReader
	clock  xtime.Clocker
	ctx    iface.ContextReader
	token  iface.TokenGenerateValidator
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

func (t *Token) ValidateRefreshToken(token *string) error {
	if err := t.token.Validate(token); err != nil {
		return xerr.InvalidToken
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
