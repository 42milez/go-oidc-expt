package service

import (
	"context"

	"github.com/42milez/go-oidc-server/cmd/entity"
	"github.com/42milez/go-oidc-server/cmd/security"
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
	hash, err := security.HashPassword(pw)

	if err != nil {
		return nil, err
	}

	ret, err := cu.repo.CreateUser(ctx, name, hash)

	if err != nil {
		return nil, err
	}

	return ret, nil
}
