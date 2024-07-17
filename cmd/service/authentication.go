package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/repository"
	"github.com/42milez/go-oidc-expt/cmd/security"
	"github.com/42milez/go-oidc-expt/pkg/xerr"
)

func NewAuthenticate(opt *option.Option) *Authenticate {
	return &Authenticate{
		repo:  repository.NewUser(opt),
		token: opt.Token,
	}
}

type Authenticate struct {
	repo  UserConsentReader
	token iface.TokenGenerator
}

func (a *Authenticate) VerifyConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) (bool, error) {
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

	ok, err := security.ComparePassword(pw, user.Password())
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, xerr.PasswordNotMatched
	}

	return user.ID(), nil
}
