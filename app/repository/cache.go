package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

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
	ret, err := s.cache.Client.HGet(ctx, key, field).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", xerr.CacheKeyNotFound
		} else {
			return "", xerr.UnexpectedErrorOccurred
		}
	}
	return ret, nil
}

func (s *Cache) ReadHashAll(ctx context.Context, key string) (map[string]string, error) {
	ret, err := s.cache.Client.HGetAll(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, xerr.CacheKeyNotFound
		} else {
			return nil, xerr.UnexpectedErrorOccurred
		}
	}
	return ret, nil
}

func (s *Cache) Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	return s.cache.Client.SetNX(ctx, key, value, ttl).Result()
}

func (s *Cache) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) (bool, error) {
	_, err := s.cache.Client.HSet(ctx, key, values).Result()
	if err != nil {
		return false, err
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
