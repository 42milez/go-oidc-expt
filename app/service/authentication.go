package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/ent/ent"
)

var errEntNotFoundError = &ent.NotFoundError{}

type Authenticate struct {
	repo     UserSelector
	tokenGen TokenGenerator
}

func (p *Authenticate) Authenticate(ctx context.Context, name, pw string) (typedef.UserID, error) {
	user, err := p.repo.SelectByName(ctx, name)

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
