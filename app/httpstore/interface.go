package httpstore

import (
	"context"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/typedef"
)

type SessionCreator interface {
	Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error)
}

type SessionReader interface {
	Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error)
}

type SessionUpdater interface {
	Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error)
}
