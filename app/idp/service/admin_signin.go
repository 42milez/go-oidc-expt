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
	errFailedToGenerateToken Err = "failed to generate token"
	errFailedToSelectAdmin   Err = "failed to select admin"
	errInvalidPassword       Err = "invalid password"
)

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
