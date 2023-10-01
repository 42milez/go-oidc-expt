package service

import (
	"context"
	"time"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewToken(db *datastore.Database) *Token {
	return &Token{
		authCodeRepo:    repository.NewAuthCode(db),
		redirectUriRepo: repository.NewRedirectUri(db),
	}
}

type Token struct {
	authCodeRepo    AuthCodeReader
	redirectUriRepo RedirectUriReader
}

func (t *Token) ValidateAuthCode(ctx context.Context, code, clientId string) error {
	authCode, err := t.authCodeRepo.ReadAuthCode(ctx, code, clientId)
	if err != nil {
		return err
	}

	if authCode.ExpireAt.After(time.Now()) {
		return xerr.AuthCodeExpired
	}

	if authCode.UsedAt != nil {
		return xerr.AuthCodeUsed
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
