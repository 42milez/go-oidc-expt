package service

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/pkg/xutil"
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

type AdminSignUp struct {
	Repo           IdentityCreator[ent.Admin]
	TokenGenerator TokenGenerator
}

func (p *AdminSignUp) SignUp(ctx context.Context, name, pw string) (*ent.Admin, error) {
	hash, err := xutil.GeneratePasswordHash(pw)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errFailedToGeneratePassword, err)
	}

	admin, err := p.Repo.Create(ctx, name, hash)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errFailedToCreateAdmin, err)
	}

	return admin, nil
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

	if err = xutil.ComparePassword(admin.Password, pw); err != nil {
		return "", fmt.Errorf("%w: %w", errInvalidPassword, err)
	}

	token, err := p.TokenGenerator.GenerateAccessToken(admin.Name)

	if err != nil {
		return "", fmt.Errorf("%w: %w", errFailedToGenerateToken, err)
	}

	return string(token), nil
}
