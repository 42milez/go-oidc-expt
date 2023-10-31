package httpstore

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
)

const nRetryWriteSession = 3
const clientIdFieldName = "ClientId"
const redirectUriFieldName = "RedirectUri"
const userIdFieldName = "UserId"

func NewReadSession(repo SessionReader) *ReadSession {
	return &ReadSession{
		repo: repo,
	}
}

type ReadSession struct {
	repo SessionReader
}

func (rs *ReadSession) ReadAuthParam(ctx context.Context, clientId, authCode string) (*typedef.AuthParam, error) {
	key := authParamSessionKey(clientId, authCode)

	redirectUri, err := rs.repo.ReadHash(ctx, key, redirectUriFieldName)
	if err != nil {
		return nil, err
	}

	userId, err := rs.repo.ReadHash(ctx, key, userIdFieldName)
	if err != nil {
		return nil, err
	}
	userIdUint64, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	return &typedef.AuthParam{
		RedirectUri: redirectUri,
		UserId:      typedef.UserID(userIdUint64),
	}, nil
}

func (rs *ReadSession) ReadRefreshTokenPermission(ctx context.Context, token string) (*typedef.AuthParam, error) {
	key := refreshTokenPermissionSessionKey(token)

	redirectUri, err := rs.repo.ReadHash(ctx, key, clientIdFieldName)
	if err != nil {
		return nil, err
	}

	userId, err := rs.repo.ReadHash(ctx, key, userIdFieldName)
	if err != nil {
		return nil, err
	}
	userIdUint64, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	return &typedef.AuthParam{
		RedirectUri: redirectUri,
		UserId:      typedef.UserID(userIdUint64),
	}, nil
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

	key := userInfoSessionKey(sid)
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

func (ws *WriteSession) WriteAuthParam(ctx context.Context, param *typedef.AuthParam, clientId, authCode string) error {
	key := authParamSessionKey(clientId, authCode)
	values := map[string]string{
		redirectUriFieldName: param.RedirectUri,
		userIdFieldName:      strconv.FormatUint(uint64(param.UserId), 10),
	}
	ok, err := ws.repo.WriteHash(ctx, key, values, config.SessionTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.CacheKeyDuplicated
	}
	return nil
}

func (ws *WriteSession) WriteRefreshTokenPermission(ctx context.Context, token, clientId string, uid typedef.UserID) error {
	key := refreshTokenPermissionSessionKey(token)
	values := map[string]string{
		clientIdFieldName: clientId,
		userIdFieldName:   strconv.FormatUint(uint64(uid), 10),
	}
	ok, err := ws.repo.WriteHash(ctx, key, values, config.SessionTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.CacheKeyDuplicated
	}

	return nil
}

func (ws *WriteSession) WriteUserInfo(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error) {
	var sid uint64
	var ok bool
	var err error

	for i := 0; i < nRetryWriteSession; i++ {
		if sid, err = ws.idGen.NextID(); err != nil {
			return 0, err
		}
		key := userInfoSessionKey(typedef.SessionID(sid))
		if ok, err = ws.repo.Write(ctx, key, uid, config.SessionTTL); err != nil {
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

func authParamSessionKey(clientId, authCode string) string {
	return fmt.Sprintf("auth:param:%s.%s", clientId, authCode)
}

func refreshTokenPermissionSessionKey(token string) string {
	return fmt.Sprintf("rp:refreshtoken:permission:%s", token)
}

func userInfoSessionKey(sid typedef.SessionID) string {
	return fmt.Sprintf("idp:session:%d", sid)
}
