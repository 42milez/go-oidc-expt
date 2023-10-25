package httpstore

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
	"github.com/google/uuid"
)

const nRetrySaveSession = 3

func NewCreateSession(repo SessionCreator) *CreateSession {
	return &CreateSession{
		repo: repo,
	}
}

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

func NewReadSession(repo SessionReader) *ReadSession {
	return &ReadSession{
		repo: repo,
	}
}

type ReadSession struct {
	repo SessionReader
}

func (rs *ReadSession) Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error) {
	return rs.repo.Read(ctx, sid)
}

func NewRestoreSession(repo SessionReader) *RestoreSession {
	return &RestoreSession{
		repo: repo,
	}
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
	ctx = context.WithValue(ctx, typedef.SessionIDKey{}, sid)
	ctx = context.WithValue(ctx, typedef.SessionKey{}, sess)
	ctx = context.WithValue(ctx, typedef.UserIDKey{}, sess.UserID)

	return r.Clone(ctx), nil
}

func NewUpdateSession(repo SessionUpdater) *UpdateSession {
	return &UpdateSession{
		repo: repo,
	}
}

type UpdateSession struct {
	repo SessionUpdater
}

func (us *UpdateSession) Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) error {
	_, err := us.repo.Update(ctx, sid, sess)
	if err != nil {
		return err
	}
	return nil
}
