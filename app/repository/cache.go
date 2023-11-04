package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/datastore"
)

func NewCache(cache *datastore.Cache) *Cache {
	return &Cache{
		cache: cache,
	}
}

type Cache struct {
	cache *datastore.Cache
}

func (s *Cache) Read(ctx context.Context, key string) (string, error) {
	v, err := s.cache.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	if xutil.IsEmpty(v) {
		return "", xerr.CacheKeyNotFound
	}
	return v, nil
}

func (s *Cache) ReadHash(ctx context.Context, key string, field string) (string, error) {
	return s.cache.Client.HGet(ctx, key, field).Result()
}

func (s *Cache) ReadHashAll(ctx context.Context, key string) (map[string]string, error) {
	return s.cache.Client.HGetAll(ctx, key).Result()
}

func (s *Cache) Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	return s.cache.Client.SetNX(ctx, key, value, ttl).Result()
}

func (s *Cache) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) (bool, error) {
	for field, value := range values {
		ok, err := s.cache.Client.HSetNX(ctx, key, field, value).Result()
		if err != nil {
			return false, err
		}
		if !ok {
			return false, fmt.Errorf("field already exists: %s", field)
		}
	}
	ok, err := s.cache.Client.Expire(ctx, key, ttl).Result()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, xerr.FailedToSetTimeoutOnCacheKey
	}
	return true, nil
}
