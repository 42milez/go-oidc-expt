package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/rs/zerolog/log"
)

type Authenticate struct {
	Repo           UserSelector
	TokenGenerator TokenGenerator
}

func (p *Authenticate) Authenticate(ctx context.Context, name, pw string) (string, error) {
	user, err := p.Repo.SelectByName(ctx, name)

	if err != nil {
		log.Error().Err(err).Msg(xerr.Wrap(errFailedToSelectUser, err).Error())
		return "", err
	}

	ok, err := auth.ComparePassword(pw, user.PasswordHash)

	if err != nil {
		log.Error().Err(err).Msg(xerr.Wrap(errPasswordNotMatched, err).Error())
		return "", err
	}

	if !ok {
		log.Error().Err(err).Msg(xerr.Wrap(errInvalidPassword, err).Error())
		return "", err
	}

	token, err := p.TokenGenerator.GenerateAccessToken(user.Name)

	if err != nil {
		log.Error().Err(err).Msg(xerr.Wrap(errFailedToGenerateToken, err).Error())
		return "", err
	}

	return string(token), nil
}
