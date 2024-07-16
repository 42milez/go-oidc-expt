package httpstore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/config"
	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/repository"
	"github.com/42milez/go-oidc-expt/pkg/xerr"
)

const redirectURIFieldName = "RedirectURI"
const nonceFieldName = "Nonce"
const userIDFieldName = "UserID"
const authTimeFieldName = "AuthTime"

func NewCache(opt *option.Option) *Cache {
	return &Cache{
		repo:  repository.NewCache(opt.Cache),
		idGen: opt.IDGen,
		token: opt.Token,
	}
}

type Cache struct {
	repo  CacheReadWriter
	idGen iface.IDGenerator
	token iface.TokenParser
}

func (c *Cache) ReadAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string) (*typedef.AuthorizationRequestFingerprint, error) {
	key := authorizationFingerprintCacheKey(clientID, authCode)
	values, err := c.repo.ReadHashAll(ctx, key)
	if errors.Is(err, xerr.CacheKeyNotFound) {
		return nil, xerr.UnauthorizedRequest
	}

	userID, err := strconv.ParseUint(values[userIDFieldName], 10, 64)
	if err != nil {
		return nil, err
	}
	authTimeUnix, err := strconv.ParseInt(values[authTimeFieldName], 10, 64)
	if err != nil {
		return nil, err
	}
	authTime := time.Unix(authTimeUnix, 0)

	return &typedef.AuthorizationRequestFingerprint{
		RedirectURI: values[redirectURIFieldName],
		Nonce:       values[nonceFieldName],
		UserID:      typedef.UserID(userID),
		AuthTime:    authTime,
	}, nil
}

type Session struct {
	ID       typedef.SessionID
	UserID   typedef.UserID
	AuthTime time.Time
}
type SessionKey struct{}

func (c *Cache) Restore(r *http.Request, sid typedef.SessionID) (*http.Request, error) {
	ctx := r.Context()

	key := sessionCacheKey(sid)
	sess := &Session{
		ID: sid,
	}

	uidRaw, err := c.repo.ReadHash(ctx, key, userIDFieldName)
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

	ctx = context.WithValue(ctx, SessionKey{}, sess)

	return r.Clone(ctx), nil
}

func (c *Cache) WriteAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string, param *typedef.AuthorizationRequestFingerprint) error {
	key := authorizationFingerprintCacheKey(clientID, authCode)
	values := map[string]any{
		redirectURIFieldName: param.RedirectURI,
		nonceFieldName:       param.Nonce,
		userIDFieldName:      strconv.FormatUint(uint64(param.UserID), 10),
		authTimeFieldName:    param.AuthTime.Unix(),
	}
	if err := c.repo.WriteHash(ctx, key, values, config.AuthCodeTTL); err != nil {
		return err
	}
	return nil
}

func (c *Cache) WriteRefreshToken(ctx context.Context, token string, clientID typedef.ClientID, userID typedef.UserID) error {
	key := refreshTokenCacheKey(clientID, userID)
	if err := c.repo.Write(ctx, key, token, config.RefreshTokenTTL); err != nil {
		return err
	}
	return nil
}

func (c *Cache) ReadRefreshToken(ctx context.Context, clientID typedef.ClientID, userID typedef.UserID) (jwt.Token, error) {
	key := refreshTokenCacheKey(clientID, userID)
	v, err := c.repo.Read(ctx, key)
	if err != nil {
		return nil, err
	}
	token, err := c.token.Parse(v)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (c *Cache) CreateSession(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error) {
	sid, err := c.idGen.NextID()
	if err != nil {
		return 0, err
	}

	key := sessionCacheKey(typedef.SessionID(sid))
	values := map[string]any{
		userIDFieldName:   strconv.FormatUint(uint64(uid), 10),
		authTimeFieldName: time.Now().Unix(),
	}
	if err = c.repo.WriteHash(ctx, key, values, config.SessionTTL); err != nil {
		return 0, err
	}

	return typedef.SessionID(sid), nil
}

func authorizationFingerprintCacheKey(clientID typedef.ClientID, authCode string) string {
	return fmt.Sprintf("op:authorization:fingerprint:%s.%s", clientID, authCode)
}

func refreshTokenCacheKey(clientID typedef.ClientID, userID typedef.UserID) string {
	return fmt.Sprintf("rp:refreshtoken:%s.%s", clientID, userID)
}

func sessionCacheKey(sid typedef.SessionID) string {
	return fmt.Sprintf("op:session:%d", sid)
}
