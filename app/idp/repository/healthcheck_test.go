package repository

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestCheckHealth_PingCache(t *testing.T) {
	ch := &CheckHealth{
		cache: xtestutil.NewCache(t),
	}
	if err := ch.PingCache(context.Background()); err != nil {
		t.Error(err)
	}
}

func TestCheckHealth_PingDB(t *testing.T) {
	ch := &CheckHealth{
		db: xtestutil.NewDatabase(t, nil),
	}
	if err := ch.PingDatabase(context.Background()); err != nil {
		t.Error(err)
	}
}
