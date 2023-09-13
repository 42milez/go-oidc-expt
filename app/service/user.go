package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
)

type CreateUser struct {
	repo UserCreator
}

func (cu *CreateUser) CreateUser(ctx context.Context, name, pw string) (*ent.User, error) {
	hash, err := xargon2.HashPassword(pw)

	if err != nil {
		return nil, err
	}

	ret, err := cu.repo.CreateUser(ctx, name, hash)

	if err != nil {
		return nil, err
	}

	return ret, nil
}
