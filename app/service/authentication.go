package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/ent/ent"

	"github.com/42milez/go-oidc-server/pkg/xargon2"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/ent/typedef"
)

var errEntNotFoundError = &ent.NotFoundError{}

type Authenticate struct {
	Repo  UserSelector
	Token TokenGenerator
}

func (p *Authenticate) Authenticate(ctx context.Context, name, pw string) (typedef.UserID, error) {
	user, err := p.Repo.SelectByName(ctx, name)

	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return 0, xerr.UserNotFound.Wrap(err)
		} else {
			return 0, err
		}
	}

	ok, err := xargon2.ComparePassword(pw, user.Password)

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, xerr.PasswordNotMatched
	}

	return user.ID, nil
}
