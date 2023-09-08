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

type IDKey struct{}
type UserIDKey struct{}

type CreateSession struct {
	repo SessionCreator
}

func (p *CreateSession) Create(ctx context.Context, sess *entity.Session) (string, error) {
	var id uuid.UUID
	var ok bool
	var err error

	for i := 0; i < nRetrySaveSession; i++ {
		if id, err = uuid.NewRandom(); err != nil {
			return "", err
		}
		if ok, err = p.repo.Create(ctx, typedef.SessionID(id.String()), sess); err != nil {
			return "", err
		}
		if ok {
			break
		}
	}

	if !ok {
		return "", xerr.SessionIDAlreadyExists
	}

	return id.String(), nil
}

type RestoreSession struct {
	repo SessionReader
}

func (p *RestoreSession) Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error) {
	sess, err := p.repo.Read(r.Context(), sid)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, IDKey{}, sid)
	ctx = context.WithValue(ctx, UserIDKey{}, sess.UserID)

	return r.Clone(ctx), nil
}

type UpdateSession struct {
	Repo SessionUpdater
}

func (p *UpdateSession) Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) error {
	_, err := p.Repo.Update(ctx, sid, sess)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionID(ctx context.Context) string {
	return ctx.Value(IDKey{}).(string)
}

func GetUserID(ctx context.Context) (typedef.UserID, bool) {
	ret, ok := ctx.Value(UserIDKey{}).(typedef.UserID)
	return ret, ok
}
