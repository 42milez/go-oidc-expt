package repository

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type CheckHealth struct {
	Cache *redis.Client
	DB    *sql.DB
}

func (p *CheckHealth) PingCache(ctx context.Context) error {
	if err := p.Cache.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (p *CheckHealth) PingDB(ctx context.Context) error {
	if err := p.DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
