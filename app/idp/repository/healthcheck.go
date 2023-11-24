package repository

import (
	"context"

	datastore2 "github.com/42milez/go-oidc-server/app/idp/datastore"
)

func NewCheckHealth(db *datastore2.Database, cache *datastore2.Cache) *CheckHealth {
	return &CheckHealth{
		db:    db,
		cache: cache,
	}
}

type CheckHealth struct {
	cache *datastore2.Cache
	db    *datastore2.Database
}

func (ch *CheckHealth) PingCache(ctx context.Context) error {
	return ch.cache.Ping(ctx)
}

func (ch *CheckHealth) PingDatabase(ctx context.Context) error {
	return ch.db.Ping(ctx)
}
