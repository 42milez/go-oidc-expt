package repository

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestHealthCheck_PingCache(t *testing.T) {
	cache := xtestutil.NewCache(t)
	ch := &CheckHealth{
		cache: cache,
	}
	if err := ch.PingCache(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingCache, err)
	}
}

func TestHealthCheck_PingDB(t *testing.T) {
	db := xtestutil.NewDatabase(t)
	ch := &CheckHealth{
		db: db,
	}
	if err := ch.PingDB(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingDatabase, err)
	}
}
