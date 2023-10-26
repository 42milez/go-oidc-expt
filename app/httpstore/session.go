package httpstore

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
	"net/http"
	"strconv"
)

const nRetryWriteSession = 3

func userIdSessionKeySchema(sid typedef.SessionID) string {
	return fmt.Sprintf("session:userid:%d", sid)
}

func redirectUriSessionKeySchema(sid typedef.SessionID) string {
	return fmt.Sprintf("session:redirecturi:%d", sid)
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

func (ss *WriteSession) SaveUserId(ctx context.Context, userId typedef.UserID) (typedef.SessionID, error) {
	var sid uint64
	var ok bool
	var err error

	for i := 0; i < nRetryWriteSession; i++ {
		if sid, err = ss.idGen.NextID(); err != nil {
			return 0, err
		}
		key := userIdSessionKeySchema(typedef.SessionID(sid))
		if ok, err = ss.repo.Write(ctx, key, userId, config.SessionTTL); err != nil {
			return 0, err
		}
		if ok {
			break
		}
	}

	if !ok {
		return 0, xerr.FailedToSaveSession
	}

	return typedef.SessionID(sid), nil
}

func (ss *WriteSession) SaveRedirectUri(ctx context.Context, sid typedef.SessionID, uri string) error {
	key := redirectUriSessionKeySchema(sid)
	ok, err := ss.repo.Write(ctx, key, uri, config.SessionTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.FailedToSaveSession
	}
	return nil
}

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

	ctx = context.WithValue(ctx, typedef.SessionIdKey{}, sid)

	uidUint64, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, typedef.UserIdKey{}, typedef.UserID(uidUint64))

	return r.Clone(ctx), nil
}
