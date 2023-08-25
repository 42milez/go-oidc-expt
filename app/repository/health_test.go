package repository

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/xtestutil"
)

func TestHealthCheck_PingCache(t *testing.T) {
	client := xtestutil.OpenRedis(t)
	ch := &CheckHealth{
		Cache: client,
	}
	if err := ch.PingCache(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingCache, err)
	}
}

func TestHealthCheck_PingDB(t *testing.T) {
	_, client := xtestutil.OpenDB(t)
	ch := &CheckHealth{
		DB: client,
	}
	if err := ch.PingDB(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingDatabase, err)
	}
}
