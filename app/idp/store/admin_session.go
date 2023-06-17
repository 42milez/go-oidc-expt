package store

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/pkg/util"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

type SessionErr uint

const (
	ErrNotFound SessionErr = iota
)

func (v SessionErr) Error() string {
	return []string{
		"not found",
	}[v]
}

func NewAdminSession(ctx context.Context, cfg *config.Config) (*AdminSession, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &AdminSession{
		client: client,
	}, nil
}

type AdminSession struct {
	client *redis.Client
}

func (p *AdminSession) Close() error {
	return p.client.Close()
}

func (p *AdminSession) SaveID(ctx context.Context, key string, id alias.AdminID) error {
	return p.client.Set(ctx, key, uint64(id), sessionTTL).Err()
}

func (p *AdminSession) LoadID(ctx context.Context, key string) (alias.AdminID, error) {
	id, err := p.client.Get(ctx, key).Uint64()
	if err != nil {
		return 0, fmt.Errorf("failed to load %q: %w", key, ErrNotFound)
	}
	return alias.AdminID(id), nil
}

func (p *AdminSession) DeleteID(ctx context.Context, key string) error {
	return p.client.Del(ctx, key).Err()
}

func (p *AdminSession) ExtractToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	j, err := auth.NewJWT(util.RealClocker{})

	if err != nil {
		return nil, fmt.Errorf("failed to create jwt: %w", err)
	}

	token, err := j.ParseRequest(r)

	if err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	if err := j.Validate(token); err != nil {
		return nil, fmt.Errorf("invalid token %q: %w", token, err)
	}

	if _, err := p.LoadID(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("%q already expired: %w", token.JwtID(), err)
	}

	return token, nil
}
