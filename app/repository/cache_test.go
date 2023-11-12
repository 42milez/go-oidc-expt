package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
)

func TestCache_Read(t *testing.T) {
	t.Parallel()

	repo := Cache{
		cache: xtestutil.NewCache(t),
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestCache_Read_OK"
		wantUserID := "486937312744178029"

		if err := repo.cache.Client.SetNX(ctx, sid, wantUserID, config.SessionTTL).Err(); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, sid)
		})

		gotUserID, err := repo.Read(ctx, sid)
		if err != nil {
			t.Fatal(err)
		}

		if wantUserID != gotUserID {
			t.Fatalf("want = %s; got = %s", wantUserID, gotUserID)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestCache_Read_NotFound"

		_, err := repo.Read(ctx, sid)
		if err == nil || !errors.Is(err, xerr.CacheKeyNotFound) {
			t.Fatalf("want = %+v; got = %+v", xerr.CacheKeyNotFound, err)
		}
	})
}

func TestCache_Write(t *testing.T) {
	t.Parallel()

	repo := Cache{
		cache: xtestutil.NewCache(t),
	}

	ctx := context.Background()
	sid := "TestCache_Write"
	uid := typedef.UserID(475924035230777348)

	t.Cleanup(func() {
		repo.cache.Client.Del(ctx, sid)
	})

	ok, err := repo.Write(ctx, sid, uid, config.SessionTTL)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal(xerr.FailedToWriteCache)
	}
}
