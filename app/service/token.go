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
		authCodeRepo: repository.NewAuthCode(db),
	}
}

type Token struct {
	authCodeRepo AuthCodeReader
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

func (t *Token) ValidateRedirectUri(ctx context.Context, uri string) error {
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
