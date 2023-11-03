package httpstore

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/iface"
	"github.com/42milez/go-oidc-server/app/option"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
)

const nRetryWriteCache = 3
const redirectUriFieldName = "RedirectUri"
const userIdFieldName = "UserId"

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

	redirectUri, err := c.repo.ReadHash(ctx, key, redirectUriFieldName)
	if err != nil {
		return nil, err
	}

	userId, err := c.repo.ReadHash(ctx, key, userIdFieldName)
	if err != nil {
		return nil, err
	}
	userIdUint64, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	return &typedef.OpenIdParam{
		RedirectUri: redirectUri,
		UserId:      typedef.UserID(userIdUint64),
	}, nil
}

func (c *Cache) ReadRefreshTokenOwner(ctx context.Context, token string) (string, error) {
	key := refreshTokenPermissionCacheKey(token)
	return c.repo.Read(ctx, key)
}

func (c *Cache) Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error) {
	ctx := r.Context()

	key := userInfoCacheKey(sid)
	uid, err := c.repo.ReadHash(ctx, key, userIdFieldName)
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

func (c *Cache) WriteOpenIdParam(ctx context.Context, param *typedef.OpenIdParam, clientId, authCode string) error {
	key := openIdParamCacheKey(clientId, authCode)
	values := map[string]string{
		redirectUriFieldName: param.RedirectUri,
		userIdFieldName:      strconv.FormatUint(uint64(param.UserId), 10),
	}
	ok, err := c.repo.WriteHash(ctx, key, values, config.CacheTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.CacheKeyDuplicated
	}
	return nil
}

func (c *Cache) WriteRefreshTokenOwner(ctx context.Context, token, clientId string) error {
	key := refreshTokenPermissionCacheKey(token)
	ok, err := c.repo.Write(ctx, key, clientId, config.CacheTTL)
	if err != nil {
		return err
	}
	if !ok {
		return xerr.CacheKeyDuplicated
	}
	return nil
}

func (c *Cache) WriteUserInfo(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error) {
	var sid uint64
	var ok bool
	var err error

	for i := 0; i < nRetryWriteCache; i++ {
		if sid, err = c.idGen.NextID(); err != nil {
			return 0, err
		}
		key := userInfoCacheKey(typedef.SessionID(sid))
		values := map[string]string{
			userIdFieldName: strconv.FormatUint(uint64(uid), 10),
		}
		if ok, err = c.repo.WriteHash(ctx, key, values, config.CacheTTL); err != nil {
			return 0, err
		}
		if ok {
			break
		}
	}

	if !ok {
		return 0, xerr.FailedToWriteCache
	}

	return typedef.SessionID(sid), nil
}

func openIdParamCacheKey(clientId, authCode string) string {
	return fmt.Sprintf("idp:openid:param:%s.%s", clientId, authCode)
}

func refreshTokenPermissionCacheKey(token string) string {
	return fmt.Sprintf("rp:refreshtoken:permission:%s", token)
}

func userInfoCacheKey(sid typedef.SessionID) string {
	return fmt.Sprintf("idp:session:%d", sid)
}