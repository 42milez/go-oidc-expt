package service

import (
	"context"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/app/idp/auth"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

const (
	errFailedToCreateAdmin      xerr.Err = "failed to create admin"
	errFailedToGeneratePassword xerr.Err = "failed to generate password"
	errFailedToGenerateToken    xerr.Err = "failed to generate token"
	errFailedToSelectAdmin      xerr.Err = "failed to select admin"
	errInvalidPassword          xerr.Err = "invalid password"
	errPasswordNotMatched       xerr.Err = "password not matched"
)

type AdminCreator struct {
	Repo IdentityCreator[ent.Admin]
}

func (p *AdminCreator) Create(ctx context.Context, name, pw string) (*ent.Admin, error) {
	hash, err := auth.GeneratePasswordHash(pw)

	if err != nil {
		log.Error().Err(err).Msg(xerr.WrapErr(errFailedToGeneratePassword, err).Error())
		return nil, err
	}

	ret, err := p.Repo.Create(ctx, name, hash)

	if err != nil {
		log.Error().Err(err).Msg(xerr.WrapErr(errFailedToCreateAdmin, err).Error())
		return nil, err
	}

	return ret, nil
}

type AdminSignIn struct {
	Repo           IdentitySelector[ent.Admin]
	TokenGenerator TokenGenerator
}

func (p *AdminSignIn) SignIn(ctx context.Context, name, pw string) (string, error) {
	admin, err := p.Repo.SelectByName(ctx, name)

	if err != nil {
		log.Error().Err(err).Msg(xerr.WrapErr(errFailedToSelectAdmin, err).Error())
		return "", err
	}

	ok, err := auth.ComparePassword(pw, admin.PasswordHash)

	if err != nil {
		log.Error().Err(err).Msg(xerr.WrapErr(errPasswordNotMatched, err).Error())
		return "", err
	}

	if !ok {
		log.Error().Err(err).Msg(xerr.WrapErr(errInvalidPassword, err).Error())
		return "", err
	}

	token, err := p.TokenGenerator.GenerateAccessToken(admin.Name)

	if err != nil {
		log.Error().Err(err).Msg(xerr.WrapErr(errFailedToGenerateToken, err).Error())
		return "", err
	}

	return string(token), nil
}
