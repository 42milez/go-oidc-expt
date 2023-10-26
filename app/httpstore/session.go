package httpstore

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/iface"

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

func (rs *ReadSession) ReadRedirectUri(ctx context.Context, clientId, authCode string) (string, error) {
	key := redirectUriSessionKey(clientId, authCode)
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

	key := userIdSessionKey(sid)
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

func NewWriteSession(repo SessionWriter, ctx iface.ContextReader, idGen IdGenerator) *WriteSession {
	return &WriteSession{
		repo:  repo,
		ctx:   ctx,
		idGen: idGen,
	}
}

type WriteSession struct {
	repo  SessionWriter
	ctx   iface.ContextReader
	idGen IdGenerator
}

func (ws *WriteSession) WriteRedirectUriAssociation(ctx context.Context, uri, clientId, authCode string) error {
	key := redirectUriSessionKey(clientId, authCode)
	ok, err := ws.repo.Write(ctx, key, uri, config.SessionTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.FailedToWriteSession
	}

	return nil
}

func (ws *WriteSession) WriteRefreshTokenOwner(ctx context.Context, token, clientId string) error {
	hash := sha256.Sum256([]byte(token))
	key := refreshTokenSessionKey(string(hash[:]))
	ok, err := ws.repo.Write(ctx, key, clientId, config.SessionTTL)
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
		key := userIdSessionKey(typedef.SessionID(sid))
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

func redirectUriSessionKey(clientId, authCode string) string {
	return fmt.Sprintf("session:redirecturi:%s.%s", clientId, authCode)
}

func refreshTokenSessionKey(token string) string {
	return fmt.Sprintf("session:refreshtoken:%s", token)
}

func userIdSessionKey(sid typedef.SessionID) string {
	return fmt.Sprintf("session:userid:%d", sid)
}
