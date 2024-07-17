package repository

import (
	"context"

	"github.com/42milez/go-oidc-expt/cmd/datastore"
)

func NewCheckHealth(db *datastore.Database, cache *datastore.Cache) *CheckHealth {
	return &CheckHealth{
		db:    db,
		cache: cache,
	}
}

type CheckHealth struct {
	cache *datastore.Cache
	db    *datastore.Database
}

func (ch *CheckHealth) PingCache(ctx context.Context) error {
	return ch.cache.Ping(ctx)
}

func (ch *CheckHealth) PingDatabase(ctx context.Context) error {
	return ch.db.Ping(ctx)
}
