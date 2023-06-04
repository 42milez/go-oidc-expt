package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/entity"
)

type CreateUser struct {
	DB   *ent.Client
	Repo UserCreater
}

func (p *CreateUser) Create(ctx context.Context) error {
	return nil
}

type ReadUser struct {
	DB   *ent.Client
	Repo UserReader
}

func (p *ReadUser) ReadUser(ctx context.Context) (entity.User, error) {
	return nil, errors.New("not implemented")
}

type ReadUsers struct {
	DB   *ent.Client
	Repo UsersReader
}

func (p *ReadUsers) ReadUsers(ctx context.Context) (entity.Users, error) {
	return nil, errors.New("not implemented")
}
