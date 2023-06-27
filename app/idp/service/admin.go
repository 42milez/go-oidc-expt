package service

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/idp/auth"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type Err string

func (v Err) Error() string {
	return string(v)
}

const (
	errFailedToCreateAdmin      Err = "failed to create admin"
	errFailedToGeneratePassword Err = "failed to generate password"
	errFailedToGenerateToken    Err = "failed to generate token"
	errFailedToSelectAdmin      Err = "failed to select admin"
	errInvalidPassword          Err = "invalid password"
)

type AdminCreator struct {
	Repo IdentityCreator[ent.Admin]
}

func (p *AdminCreator) Create(ctx context.Context, name, pw string) (*ent.Admin, error) {
	hash, err := auth.GeneratePasswordHash(pw)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errFailedToGeneratePassword, err)
	}

	ret, err := p.Repo.Create(ctx, name, hash)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errFailedToCreateAdmin, err)
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
		return "", fmt.Errorf("%w: %w", errFailedToSelectAdmin, err)
	}

	ok, err := auth.ComparePassword(pw, admin.PasswordHash)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", fmt.Errorf("%w: %w", errInvalidPassword, err)
	}

	token, err := p.TokenGenerator.GenerateAccessToken(admin.Name)

	if err != nil {
		return "", fmt.Errorf("%w: %w", errFailedToGenerateToken, err)
	}

	return string(token), nil
}
