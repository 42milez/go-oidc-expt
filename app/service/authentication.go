package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

type Authenticate struct {
	Repo  UserSelector
	Token TokenGenerator
}

func (p *Authenticate) Authenticate(ctx context.Context, name, pw string) (typedef.UserID, error) {
	user, err := p.Repo.SelectByName(ctx, name)

	if err != nil {
		return "", err
	}

	ok, err := auth.ComparePassword(pw, user.Password)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", xerr.PasswordNotMatched
	}

	return user.ID, nil
}
