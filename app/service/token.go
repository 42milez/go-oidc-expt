package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewToken(db *datastore.Database, c xtime.Clocker) *Token {
	return &Token{
		authCodeRepo:    repository.NewAuthCode(db),
		redirectUriRepo: repository.NewRedirectUri(db),
		clock:           c,
	}
}

type Token struct {
	authCodeRepo    AuthCodeReadMarker
	redirectUriRepo RedirectUriReader
	clock           xtime.Clocker
}

func (t *Token) ValidateAuthCode(ctx context.Context, code, clientId string) error {
	authCode, err := t.authCodeRepo.ReadAuthCode(ctx, code, clientId)
	if err != nil {
		return err
	}

	if !authCode.ExpireAt.After(t.clock.Now()) {
		return xerr.AuthCodeExpired
	}

	if authCode.UsedAt != nil {
		return xerr.AuthCodeUsed
	}

	return nil
}

func (t *Token) RevokeAuthCode(ctx context.Context, code, clientId string) error {
	_, err := t.authCodeRepo.MarkAuthCodeUsed(ctx, code, clientId)
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) ValidateRedirectUri(ctx context.Context, uri, clientId string) error {
	_, err := t.redirectUriRepo.ReadRedirectUri(ctx, uri, clientId)
	if err != nil {
		return err
	}

	// TODO: Compare the redirect uri with the one that passed to authorization endpoint.
	// ...

	return nil
}

func (t *Token) GenerateAccessToken() (string, error) {
	return "", nil
}

func (t *Token) GenerateRefreshToken() (string, error) {
	return "", nil
}

func (t *Token) GenerateIdToken() (string, error) {
	return "", nil
}
