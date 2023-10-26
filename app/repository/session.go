package repository

import (
	"context"
	"time"

	"github.com/42milez/go-oidc-server/app/datastore"
)

func NewSession(cache *datastore.Cache) *Session {
	return &Session{
		cache: cache,
	}
}

type Session struct {
	cache *datastore.Cache
}

func (s *Session) Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	return s.cache.Client.SetNX(ctx, key, value, ttl).Result()
}

func (s *Session) Read(ctx context.Context, key string) (string, error) {
	return s.cache.Client.Get(ctx, key).Result()
}
