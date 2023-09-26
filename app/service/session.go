package service

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/google/uuid"
)

const nRetrySaveSession = 3

type SessionIDKey struct{}
type SessionKey struct{}

type UserIDKey struct{}

type CreateSession struct {
	repo SessionCreator
}

func (cs *CreateSession) Create(ctx context.Context, sess *entity.Session) (string, error) {
	var id uuid.UUID
	var ok bool
	var err error

	for i := 0; i < nRetrySaveSession; i++ {
		if id, err = uuid.NewRandom(); err != nil {
			return "", err
		}
		if ok, err = cs.repo.Create(ctx, typedef.SessionID(id.String()), sess); err != nil {
			return "", err
		}
		if ok {
			break
		}
	}

	if !ok {
		return "", xerr.FailedToCreateSession
	}

	return id.String(), nil
}

type RestoreSession struct {
	repo SessionReader
}

func (rs *RestoreSession) Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error) {
	sess, err := rs.repo.Read(r.Context(), sid)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, SessionIDKey{}, sid)
	ctx = context.WithValue(ctx, SessionKey{}, sess)
	ctx = context.WithValue(ctx, UserIDKey{}, sess.UserID)

	return r.Clone(ctx), nil
}

type UpdateSession struct {
	repo SessionUpdater
}

func (up *UpdateSession) Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) error {
	_, err := up.repo.Update(ctx, sid, sess)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(ctx context.Context) (sess *entity.Session, ok bool) {
	sess, ok = ctx.Value(SessionKey{}).(*entity.Session)
	return
}

func GetSessionID(ctx context.Context) (sessId typedef.SessionID, ok bool) {
	sessId, ok = ctx.Value(SessionIDKey{}).(typedef.SessionID)
	return
}
