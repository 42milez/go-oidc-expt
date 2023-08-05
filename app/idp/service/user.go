package service

import (
	"context"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/app/idp/auth"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

const (
	errFailedToCreateUser       xerr.Err = "failed to create user"
	errFailedToGeneratePassword xerr.Err = "failed to generate password"
	errFailedToGenerateToken    xerr.Err = "failed to generate token"
	errFailedToSelectUser       xerr.Err = "failed to select user"
	errInvalidPassword          xerr.Err = "invalid password"
	errPasswordNotMatched       xerr.Err = "password not matched"
)

type CreateUser struct {
	Repo UserCreator
}

func (p *CreateUser) CreateUser(ctx context.Context, name, pw string) (*ent.User, error) {
	hash, err := auth.GeneratePasswordHash(pw)

	if err != nil {
		log.Error().Err(err).Msg(xerr.Wrap(errFailedToGeneratePassword, err).Error())
		return nil, err
	}

	ret, err := p.Repo.Create(ctx, name, hash)

	if err != nil {
		log.Error().Err(err).Msg(xerr.Wrap(errFailedToCreateUser, err).Error())
		return nil, err
	}

	return ret, nil
}
