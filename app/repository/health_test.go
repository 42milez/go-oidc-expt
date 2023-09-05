package repository

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestHealthCheck_PingCache(t *testing.T) {
	client := xtestutil.NewRedisClient(t)
	ch := &CheckHealth{
		Cache: client,
	}
	if err := ch.PingCache(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingCache, err)
	}
}

func TestHealthCheck_PingDB(t *testing.T) {
	_, client := xtestutil.NewEntClient(t)
	ch := &CheckHealth{
		DB: client,
	}
	if err := ch.PingDB(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingDatabase, err)
	}
}
