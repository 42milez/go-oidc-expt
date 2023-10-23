package service

import (
	"context"
	"strconv"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewToken(db *datastore.Database, c xtime.Clocker, sess SessionReader, token TokenGenerateValidator) *Token {
	return &Token{
		acRepo: repository.NewAuthCode(db),
		ruRepo: repository.NewRedirectUri(db),
		cr:     &httpstore.ReadContext{},
		sess:   sess,
		token:  token,
		clock:  c,
	}
}

type Token struct {
	acRepo AuthCodeReadMarker
	ruRepo RedirectUriReader
	cr     ContextReader
	sess   SessionReader
	token  TokenGenerateValidator
	clock  xtime.Clocker
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
	_, err := t.acRepo.MarkAuthCodeUsed(ctx, code, clientId)
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) ValidateRedirectUri(ctx context.Context, uri, clientId string) error {
	_, err := t.ruRepo.ReadRedirectUri(ctx, clientId)
	if err != nil {
		return err
	}

	sid, ok := t.cr.Read(ctx, typedef.SessionIDKey{}).(typedef.SessionID)
	if !ok {
		return xerr.ContextValueNotFound
	}

	sess, err := t.sess.Read(ctx, sid)
	if err != nil {
		return err
	}

	if sess.RedirectUri != uri {
		return xerr.RedirectUriNotMatched
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
	uidConverted := strconv.FormatUint(uint64(uid), 10)
	accessToken, err := t.token.GenerateAccessToken(uidConverted)
	if err != nil {
		return "", err
	}
	return string(accessToken), nil
}

func (t *Token) GenerateRefreshToken(uid typedef.UserID) (string, error) {
	uidConverted := strconv.FormatUint(uint64(uid), 10)
	refreshToken, err := t.token.GenerateRefreshToken(uidConverted)
	if err != nil {
		return "", err
	}
	return string(refreshToken), nil
}

func (t *Token) GenerateIdToken(uid typedef.UserID) (string, error) {
	uidConverted := strconv.FormatUint(uint64(uid), 10)
	idToken, err := t.token.GenerateIdToken(uidConverted)
	if err != nil {
		return "", err
	}
	return string(idToken), nil
}
