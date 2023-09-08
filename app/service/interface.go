package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/ent/ent"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type CachePingSender interface {
	PingCache(ctx context.Context) error
}

type DBPingSender interface {
	PingDB(ctx context.Context) error
}

type HealthChecker interface {
	CachePingSender
	DBPingSender
}

type TokenGenerator interface {
	MakeAccessToken(name string) ([]byte, error)
}

type UserCreator interface {
	Create(ctx context.Context, name string, pw string) (*ent.User, error)
}

type UserSelector interface {
	SelectByName(ctx context.Context, name string) (*ent.User, error)
}

type SessionCreator interface {
	Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error)
}

type SessionReader interface {
	Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error)
}

type SessionUpdater interface {
	Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error)
}
