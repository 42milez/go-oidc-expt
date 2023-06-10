package service

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/testutil"
)

func TestCheckHealth_PingCache(t *testing.T) {
	// NOT IMPLEMENTED
}

func TestCheckHealth_PingDB(t *testing.T) {
	_, db := testutil.OpenDB(t)
	ch := CheckHealth{
		DB: db,
	}
	if err := ch.PingDB(context.Background()); err != nil {
		t.Errorf("failed to ping database: %+v", err)
	}
}
