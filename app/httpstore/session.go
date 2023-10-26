package httpstore

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
)

const nRetryWriteSession = 3

func NewReadSession(repo SessionReader) *ReadSession {
	return &ReadSession{
		repo: repo,
	}
}

type ReadSession struct {
	repo SessionReader
}

func (rs *ReadSession) ReadRedirectUri(ctx context.Context, sid typedef.SessionID) (string, error) {
	key := redirectUriSessionKeySchema(sid)
	return rs.repo.Read(ctx, key)
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
	ctx := r.Context()

	key := userIdSessionKeySchema(sid)
	uid, err := rs.repo.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	uidUint64, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, typedef.SessionIdKey{}, sid)
	ctx = context.WithValue(ctx, typedef.UserIdKey{}, typedef.UserID(uidUint64))

	return r.Clone(ctx), nil
}

func NewWriteSession(repo SessionWriter, idGen IdGenerator) *WriteSession {
	return &WriteSession{
		repo:  repo,
		idGen: idGen,
	}
}

type WriteSession struct {
	repo  SessionWriter
	idGen IdGenerator
}

func (ws *WriteSession) WriteRedirectUri(ctx context.Context, sid typedef.SessionID, uri string) error {
	key := redirectUriSessionKeySchema(sid)
	ok, err := ws.repo.Write(ctx, key, uri, config.SessionTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.FailedToWriteSession
	}
	return nil
}

func (ws *WriteSession) WriteUserId(ctx context.Context, userId typedef.UserID) (typedef.SessionID, error) {
	var sid uint64
	var ok bool
	var err error

	for i := 0; i < nRetryWriteSession; i++ {
		if sid, err = ws.idGen.NextID(); err != nil {
			return 0, err
		}
		key := userIdSessionKeySchema(typedef.SessionID(sid))
		if ok, err = ws.repo.Write(ctx, key, userId, config.SessionTTL); err != nil {
			return 0, err
		}
		if ok {
			break
		}
	}

	if !ok {
		return 0, xerr.FailedToWriteSession
	}

	return typedef.SessionID(sid), nil
}

func userIdSessionKeySchema(sid typedef.SessionID) string {
	return fmt.Sprintf("session:userid:%d", sid)
}

func redirectUriSessionKeySchema(sid typedef.SessionID) string {
	return fmt.Sprintf("session:redirecturi:%d", sid)
}
