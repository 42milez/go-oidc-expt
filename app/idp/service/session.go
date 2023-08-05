package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/repository"
)

type Session struct {
	Repo *repository.Session
}

func (p *Session) SaveID(ctx context.Context) {

}

func (p *Session) LoadID(ctx context.Context) {

}
