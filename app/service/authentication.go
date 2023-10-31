package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/typedef"
)

func NewAuthenticate(repo UserConsentReader, tokenGen iface.TokenGenerator) *Authenticate {
	return &Authenticate{
		repo:     repo,
		tokenGen: tokenGen,
	}
}

type Authenticate struct {
	repo     UserConsentReader
	tokenGen iface.TokenGenerator
}

func (a *Authenticate) VerifyConsent(ctx context.Context, userID typedef.UserID, clientID string) (bool, error) {
	_, err := a.repo.ReadConsent(ctx, userID, clientID)
	if err != nil {
		if errors.Is(err, xerr.ConsentNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (a *Authenticate) VerifyPassword(ctx context.Context, name, pw string) (typedef.UserID, error) {
	user, err := a.repo.ReadUser(ctx, name)
	if err != nil {
		if errors.Is(err, xerr.UserNotFound) {
			return 0, xerr.UserNotFound
		} else {
			return 0, err
		}
	}

	ok, err := xargon2.ComparePassword(pw, user.Password())
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, xerr.PasswordNotMatched
	}

	return user.ID(), nil
}
