package httpstore

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

const clientIdFieldName = "ClientId"
const redirectURIFieldName = "RedirectURI"
const userIdFieldName = "UserId"
const authTimeFieldName = "AuthTime"

func NewCache(opt *option.Option) *Cache {
	return &Cache{
		repo:  repository.NewCache(opt.Cache),
		idGen: opt.IdGen,
	}
}

type Cache struct {
	repo  CacheReadWriter
	idGen iface.IdGenerator
}

func (c *Cache) ReadOpenIdParam(ctx context.Context, clientId, authCode string) (*typedef.OpenIdParam, error) {
	key := openIdParamCacheKey(clientId, authCode)
	values, err := c.repo.ReadHashAll(ctx, key)
	if errors.Is(err, xerr.CacheKeyNotFound) {
		return nil, xerr.UnauthorizedRequest
	}

	redirectURI := values[redirectURIFieldName]
	userID, err := strconv.ParseUint(values[userIdFieldName], 10, 64)
	if err != nil {
		return nil, err
	}

	return &typedef.OpenIdParam{
		RedirectURI: redirectURI,
		UserId:      typedef.UserID(userID),
	}, nil
}

func (c *Cache) ReadRefreshTokenPermission(ctx context.Context, token string) (*typedef.RefreshTokenPermission, error) {
	key := refreshTokenPermissionCacheKey(hash(token))
	perm, err := c.repo.ReadHashAll(ctx, key)
	if err != nil {
		return nil, err
	}
	userIdUint64, err := strconv.ParseUint(perm[userIdFieldName], 10, 64)
	if err != nil {
		return nil, err
	}
	return &typedef.RefreshTokenPermission{
		ClientId: perm[clientIdFieldName],
		UserId:   typedef.UserID(userIdUint64),
	}, nil
}

type Session struct {
	UserID   typedef.UserID
	AuthTime time.Time
}
type SessionKey struct{}

func (c *Cache) Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error) {
	ctx := r.Context()

	key := sessionCacheKey(sid)
	sess := &Session{}

	uidRaw, err := c.repo.ReadHash(ctx, key, userIdFieldName)
	if err != nil {
		return nil, err
	}
	uid, err := strconv.ParseUint(uidRaw, 10, 64)
	if err != nil {
		return nil, err
	}
	sess.UserID = typedef.UserID(uid)

	authTimeRaw, err := c.repo.ReadHash(ctx, key, authTimeFieldName)
	if err != nil {
		return nil, err
	}
	authTime, err := strconv.ParseInt(authTimeRaw, 10, 64)
	if err != nil {
		return nil, err
	}
	sess.AuthTime = time.Unix(authTime, 0)

	ctx = context.WithValue(ctx, typedef.SessionIdKey{}, sid)
	ctx = context.WithValue(ctx, SessionKey{}, sess)

	return r.Clone(ctx), nil
}

func (c *Cache) WriteOpenIdParam(ctx context.Context, param *typedef.OpenIdParam, clientId, authCode string) error {
	key := openIdParamCacheKey(clientId, authCode)
	values := map[string]any{
		redirectURIFieldName: param.RedirectURI,
		userIdFieldName:      strconv.FormatUint(uint64(param.UserId), 10),
	}
	if err := c.repo.WriteHash(ctx, key, values, config.AuthCodeTTL); err != nil {
		return err
	}
	return nil
}

func (c *Cache) WriteRefreshTokenPermission(ctx context.Context, token, clientId string, userId typedef.UserID) error {
	key := refreshTokenPermissionCacheKey(hash(token))
	values := map[string]any{
		clientIdFieldName: clientId,
		userIdFieldName:   strconv.FormatUint(uint64(userId), 10),
	}
	if err := c.repo.WriteHash(ctx, key, values, config.RefreshTokenTTL); err != nil {
		return err
	}
	return nil
}

func (c *Cache) CreateSession(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error) {
	sid, err := c.idGen.NextID()
	if err != nil {
		return 0, err
	}

	key := sessionCacheKey(typedef.SessionID(sid))
	values := map[string]any{
		userIdFieldName:   strconv.FormatUint(uint64(uid), 10),
		authTimeFieldName: time.Now().Unix(),
	}
	if err = c.repo.WriteHash(ctx, key, values, config.SessionTTL); err != nil {
		return 0, err
	}

	return typedef.SessionID(sid), nil
}

func openIdParamCacheKey(clientId, authCode string) string {
	return fmt.Sprintf("idp:openid:param:%s.%s", clientId, authCode)
}

func refreshTokenPermissionCacheKey(token string) string {
	return fmt.Sprintf("rp:refreshtoken:permission:%s", token)
}

func sessionCacheKey(sid typedef.SessionID) string {
	return fmt.Sprintf("idp:session:%d", sid)
}

func hash(s string) string {
	h := sha256.Sum256([]byte(s))
	return string(h[:])
}
