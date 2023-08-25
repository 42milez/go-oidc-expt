package service

import (
	"context"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/app/idp/auth"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type CreateUser struct {
	Repo UserCreator
}

func (p *CreateUser) CreateUser(ctx context.Context, name, pw string) (*ent.User, error) {
	hash, err := auth.HashPassword(pw)

	if err != nil {
		log.Error().Err(err).Msg(xerr.FailedToHashPassword.Wrap(err).Error())
		return nil, err
	}

	ret, err := p.Repo.Create(ctx, name, hash)

	if err != nil {
		log.Error().Err(err).Msg(xerr.FailedToCreateUser.Wrap(err).Error())
		return nil, err
	}

	return ret, nil
}
