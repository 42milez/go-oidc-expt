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
		key := "TestCache_Read_OK"
		wantValue := key

		if err := repo.cache.Client.SetNX(ctx, key, wantValue, config.SessionTTL).Err(); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, key)
		})

		gotUserID, err := repo.Read(ctx, key)
		if err != nil {
			t.Fatal(err)
		}

		if wantValue != gotUserID {
			t.Fatalf("want = %s; got = %s", wantValue, gotUserID)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestCache_Read_NotFound"

		_, err := repo.Read(ctx, key)
		if err == nil || !errors.Is(err, xerr.CacheKeyNotFound) {
			t.Fatalf("want = %+v; got = %+v", xerr.CacheKeyNotFound, err)
		}
	})
}

func TestCache_ReadHash(t *testing.T) {
	t.Parallel()

	repo := Cache{
		cache: xtestutil.NewCache(t),
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestCache_ReadHash_OK"
		wantValues := map[string]string{
			"value1": "TestCache_ReadHash_OK_Value1",
			"value2": "TestCache_ReadHash_OK_Value2",
		}

		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, key)
		})

		if err := repo.cache.Client.HSet(ctx, key, wantValues).Err(); err != nil {
			t.Fatal(err)
		}

		gotVal1, err := repo.ReadHash(ctx, key, "value1")
		if err != nil {
			t.Fatal(err)
		}

		if wantValues["value1"] != gotVal1 {
			t.Fatalf("want = %s; got = %s", wantValues["values1"], gotVal1)
		}

		gotVal2, err := repo.ReadHash(ctx, key, "value2")
		if err != nil {
			t.Fatal(err)
		}

		if wantValues["value2"] != gotVal2 {
			t.Fatalf("want = %s; got = %s", wantValues["values2"], gotVal2)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestCache_ReadHash_NotFound"

		_, err := repo.ReadHash(ctx, key, "value")
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
