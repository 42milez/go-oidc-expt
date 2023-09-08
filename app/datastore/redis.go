package datastore

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func (c *Cache) Ping(ctx context.Context) error {
	if err := c.Client.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}
