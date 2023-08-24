package service

import (
	"context"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/auth"
)

type Authenticate struct {
	Repo  UserSelector
	Token TokenGenerator
}

func (p *Authenticate) Authenticate(ctx context.Context, name, pw string) error {
	user, err := p.Repo.SelectByName(ctx, name)

	if err != nil {
		return err
	}

	ok, err := auth.ComparePassword(pw, user.Password)

	if err != nil {
		return err
	}

	if !ok {
		return xerr.ErrPasswordNotMatched
	}

	return nil
}
