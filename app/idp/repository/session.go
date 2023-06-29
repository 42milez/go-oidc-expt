package repository

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

const (
	ErrFailedToDeleteItem   xerr.Err = "failed to delete item"
	ErrFailedToExtractToken xerr.Err = "failed to extract token"
	ErrFailedToSaveItem     xerr.Err = "failed to save item"
	ErrFailedToLoadItem     xerr.Err = "failed to load item"
)

func NewAdminSession(ctx context.Context, cfg *config.Config) (*Session[typedef.AdminID], error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, xerr.WrapErr(xerr.FailedToReachHost, err)
	}

	jwtUtil, err := auth.NewJWTUtil(xutil.RealClocker{})

	if err != nil {
		return nil, xerr.WrapErr(xerr.FailedToInitialize, err)
	}

	return &Session[typedef.AdminID]{
		client: client,
		jwt:    jwtUtil,
	}, nil
}

type Session[T typedef.AdminID | typedef.UserID] struct {
	client *redis.Client
	jwt    *auth.JWTUtil
}

func (p *Session[T]) Close() error {
	return p.client.Close()
}

func (p *Session[T]) saveID(ctx context.Context, key string, id T) error {
	if err := p.client.Set(ctx, key, id, sessionTTL).Err(); err != nil {
		return xerr.WrapErr(fmt.Errorf("%w : key=%s, id=%s", ErrFailedToSaveItem, key, id), err)
	}
	return nil
}

func (p *Session[T]) load(ctx context.Context, key string) (T, error) {
	ret, err := p.client.Get(ctx, key).Result()
	if err != nil {
		return "", xerr.WrapErr(ErrFailedToLoadItem, err)
	}
	return T(ret), nil
}

func (p *Session[T]) delete(ctx context.Context, key string) error {
	if err := p.client.Del(ctx, key).Err(); err != nil {
		return xerr.WrapErr(fmt.Errorf("%w : key=%s", ErrFailedToDeleteItem, key), err)
	}
	return nil
}

type IDKey struct{}

func (p *Session[T]) setID(ctx context.Context, id T) context.Context {
	return context.WithValue(ctx, IDKey{}, id)
}

func (p *Session[T]) getID(ctx context.Context) (T, bool) {
	id, ok := ctx.Value(IDKey{}).(T)
	return id, ok
}

func (p *Session[T]) FillContext(r *http.Request) (*http.Request, error) {
	token, err := p.jwt.ExtractToken(r)

	if err != nil {
		return nil, xerr.WrapErr(ErrFailedToExtractToken, err)
	}

	id, err := p.load(r.Context(), token.JwtID())

	if err != nil {
		return nil, xerr.WrapErr(ErrFailedToLoadItem, err)
	}

	ctx := p.setID(r.Context(), id)

	return r.Clone(ctx), nil
}
