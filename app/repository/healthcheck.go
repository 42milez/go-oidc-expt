package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/datastore"
)

type CheckHealth struct {
	cache *datastore.Cache
	db    *datastore.Database
}

func (ch *CheckHealth) PingCache(ctx context.Context) error {
	return ch.cache.Ping(ctx)
}

func (ch *CheckHealth) PingDB(ctx context.Context) error {
	return ch.db.Ping(ctx)
}
