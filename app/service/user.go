package service

import (
	"context"

	"github.com/42milez/go-oidc-server/pkg/xargon2"

	"github.com/42milez/go-oidc-server/app/ent/ent"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/rs/zerolog/log"
)

type CreateUser struct {
	Repo UserCreator
}

func (p *CreateUser) CreateUser(ctx context.Context, name, pw string) (*ent.User, error) {
	hash, err := xargon2.HashPassword(pw)

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
