package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
)

func NewRegisterUser(repo UserCreator) *CreateUser {
	return &CreateUser{
		repo: repo,
	}
}

type CreateUser struct {
	repo UserCreator
}

func (cu *CreateUser) RegisterUser(ctx context.Context, name, pw string) (*entity.User, error) {
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
