package store

import (
	"context"
	"fmt"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

type SessionErr string

const (
	ErrFailedToDelete SessionErr = "failed to delete"
	ErrFailedToSaveID SessionErr = "failed to save id"
	ErrFailedToLoad   SessionErr = "failed to load"
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
	return &Session[alias.AdminID]{
		client: client,
	}, nil
}

type Session[T alias.AdminID | alias.UserID] struct {
	client *redis.Client
}

func (p *Session[T]) Close() error {
	return p.client.Close()
}

func (p *Session[T]) SaveID(ctx context.Context, key string, id T) error {
	if err := p.client.Set(ctx, key, uint64(id), sessionTTL).Err(); err != nil {
		return fmt.Errorf("%w ( key = %s, id = %d): %w", ErrFailedToSaveID, key, id, err)
	}
	return nil
}

func (p *Session[T]) Load(ctx context.Context, key string) (T, error) {
	id, err := p.client.Get(ctx, key).Uint64()
	if err != nil {
		return 0, fmt.Errorf("%w ( %s ): %w", ErrFailedToLoad, key, err)
	}
	return T(id), nil
}

func (p *Session[T]) Delete(ctx context.Context, key string) error {
	if err := p.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("%w ( %s ): %w", ErrFailedToDelete, key, err)
	}
	return nil
}

type IDKey struct{}

func (p *Session[T]) SetID(ctx context.Context, id T) context.Context {
	return context.WithValue(ctx, IDKey{}, id)
}

func (p *Session[T]) GetID(ctx context.Context) (T, bool) {
	id, ok := ctx.Value(IDKey{}).(T)
	return id, ok
}
