package store

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/pkg/util"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

type SessionErr string

const (
	ErrFailedToDelete SessionErr = "failed to delete"
	ErrFailedToSave SessionErr = "failed to save"
	ErrNotFound SessionErr = "not found"
)

func (v SessionErr) Error() string {
	return string(v)
}

func NewAdminSession(ctx context.Context, cfg *config.Config) (*Session[alias.AdminID], error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}
	return &Session[alias.AdminID]{
		client: client,
	}, nil
}

type Session[T alias.AdminID|alias.UserID] struct {
	client *redis.Client
}

func (p *Session[T]) Close() error {
	return p.client.Close()
}

func (p *Session[T]) SaveID(ctx context.Context, key string, id T) error {
	if err := p.client.Set(ctx, key, uint64(id), sessionTTL).Err(); err != nil {
		return ErrFailedToSave
	}
	return nil
}

func (p *Session[T]) LoadID(ctx context.Context, key string) (T, error) {
	id, err := p.client.Get(ctx, key).Uint64()
	if err != nil {
		return 0, ErrNotFound
	}
	return T(id), nil
}

func (p *Session[T]) DeleteID(ctx context.Context, key string) error {
	if err := p.client.Del(ctx, key).Err(); err != nil {
		return ErrFailedToDelete
	}
	return nil
}

func (p *Session[T]) ExtractToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	j, err := auth.NewJWT(util.RealClocker{})

	if err != nil {
		return nil, fmt.Errorf("failed to create jwt: %w", err)
	}

	token, err := j.ParseRequest(r)

	if err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	if err := j.Validate(token); err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if _, err := p.LoadID(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("%s ( %s ): %w", ErrNotFound, token.JwtID(), err)
	}

	return token, nil
}
