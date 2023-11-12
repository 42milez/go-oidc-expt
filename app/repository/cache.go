package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewCache(cache *datastore.Cache) *Cache {
	return &Cache{
		cache: cache,
	}
}

type Cache struct {
	cache *datastore.Cache
}

func (c *Cache) Read(ctx context.Context, key string) (string, error) {
	v, err := c.cache.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", xerr.CacheKeyNotFound
		} else {
			return "", xerr.UnexpectedErrorOccurred.Wrap(err)
		}
	}
	return v, nil
}

func (c *Cache) ReadHash(ctx context.Context, key string, field string) (string, error) {
	ret, err := c.cache.Client.HGet(ctx, key, field).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", xerr.CacheKeyNotFound
		} else {
			return "", xerr.UnexpectedErrorOccurred.Wrap(err)
		}
	}
	return ret, nil
}

func (c *Cache) ReadHashAll(ctx context.Context, key string) (map[string]string, error) {
	ret, err := c.cache.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, xerr.UnexpectedErrorOccurred.Wrap(err)
	}
	if len(ret) == 0 {
		return nil, xerr.CacheKeyNotFound
	}
	return ret, nil
}

func (c *Cache) Write(ctx context.Context, key string, value any, ttl time.Duration) error {
	ok, err := c.cache.Client.SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		return xerr.UnexpectedErrorOccurred.Wrap(err)
	}
	if !ok {
		return xerr.CacheKeyDuplicated
	}
	return nil
}

func (c *Cache) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) error {
	for fieldName, v := range values {
		ok, err := c.cache.Client.HSetNX(ctx, key, fieldName, v).Result()
		if err != nil {
			return xerr.UnexpectedErrorOccurred.Wrap(err)
		}
		if !ok {
			return xerr.CacheFieldDuplicated
		}
	}

	ok, err := c.cache.Client.Expire(ctx, key, ttl).Result()
	if err != nil {
		return xerr.UnexpectedErrorOccurred.Wrap(err)
	}
	if !ok {
		return xerr.FailedToSetTimeoutOnCacheKey
	}

	return nil
}
