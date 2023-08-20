package service

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/xtestutil"
)

func TestCheckHealth_PingCache(t *testing.T) {
	// NOT IMPLEMENTED
}

func TestCheckHealth_PingDB(t *testing.T) {
	_, db := xtestutil.OpenDB(t)
	ch := CheckHealth{
		DB: db,
	}
	if err := ch.PingDB(context.Background()); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToPingDatabase, err)
	}
}