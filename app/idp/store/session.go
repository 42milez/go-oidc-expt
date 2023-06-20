package store

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/pkg/util"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

type SessionErr string

const (
	ErrFailedToDelete       SessionErr = "failed to delete"
	ErrFailedToExtractToken SessionErr = "failed to extract token"
	ErrFailedToSaveID       SessionErr = "failed to save id"
	ErrFailedToLoad         SessionErr = "failed to load"
)

func (v SessionErr) Error() string {
	return string(v)
}

func NewAdminSession(ctx context.Context, cfg *config.Config) (*Session[alias.AdminID], error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", xerr.FailedToReachHost, err)
	}

	jwtUtil, err := auth.NewJWTUtil(util.RealClocker{})

	if err != nil {
		return nil, fmt.Errorf("%w", xerr.FailedToInitialize)
	}

	return &Session[alias.AdminID]{
		client: client,
		jwt:    jwtUtil,
	}, nil
}

type Session[T alias.AdminID | alias.UserID] struct {
	client *redis.Client
	jwt    *auth.JWTUtil
}

func (p *Session[T]) Close() error {
	return p.client.Close()
}

func (p *Session[T]) saveID(ctx context.Context, key string, id T) error {
	if err := p.client.Set(ctx, key, uint64(id), sessionTTL).Err(); err != nil {
		return fmt.Errorf("%w ( key = %s, id = %d): %w", ErrFailedToSaveID, key, id, err)
	}
	return nil
}

func (p *Session[T]) load(ctx context.Context, key string) (T, error) {
	id, err := p.client.Get(ctx, key).Uint64()
	if err != nil {
		return 0, fmt.Errorf("%w ( %s ): %w", ErrFailedToLoad, key, err)
	}
	return T(id), nil
}

func (p *Session[T]) delete(ctx context.Context, key string) error {
	if err := p.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("%w ( %s ): %w", ErrFailedToDelete, key, err)
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
		return nil, fmt.Errorf("%w: %w", ErrFailedToExtractToken, err)
	}

	id, err := p.load(r.Context(), token.JwtID())

	if err != nil {
		return nil, err
	}

	ctx := p.setID(r.Context(), id)

	return r.Clone(ctx), nil
}
