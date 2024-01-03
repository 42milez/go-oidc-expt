package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
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
		ttl := 1 * time.Minute

		if err := repo.cache.Client.SetNX(ctx, key, wantValue, ttl).Err(); err != nil {
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
		want := map[string]string{
			"value1": "TestCache_ReadHash_OK_Value1",
			"value2": "TestCache_ReadHash_OK_Value2",
		}

		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, key)
		})

		if err := repo.cache.Client.HSet(ctx, key, want).Err(); err != nil {
			t.Fatal(err)
		}

		got1, err := repo.ReadHash(ctx, key, "value1")
		if err != nil {
			t.Fatal(err)
		}

		if want["value1"] != got1 {
			t.Fatalf("want = %s; got = %s", want["values1"], got1)
		}

		got2, err := repo.ReadHash(ctx, key, "value2")
		if err != nil {
			t.Fatal(err)
		}

		if want["value2"] != got2 {
			t.Fatalf("want = %s; got = %s", want["values2"], got2)
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

func TestCache_ReadHashAll(t *testing.T) {
	t.Parallel()

	repo := Cache{
		cache: xtestutil.NewCache(t),
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestCache_ReadHashAll_OK"
		want := map[string]string{
			"value1": "TestCache_ReadHash_OK_Value1",
			"value2": "TestCache_ReadHash_OK_Value2",
		}

		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, key)
		})

		if err := repo.cache.Client.HSet(ctx, key, want).Err(); err != nil {
			t.Fatal(err)
		}

		got, err := repo.ReadHashAll(ctx, key)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("-want +got:\n%s", diff)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestCache_ReadHashAll_NotFound"

		_, err := repo.ReadHashAll(ctx, key)
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
	key := "TestCache_Write"
	value := "TestCache_Write_Value"
	ttl := 1 * time.Minute

	t.Cleanup(func() {
		repo.cache.Client.Del(ctx, key)
	})

	err := repo.Write(ctx, key, value, ttl)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCache_WriteHash(t *testing.T) {
	t.Parallel()

	repo := Cache{
		cache: xtestutil.NewCache(t),
	}

	ctx := context.Background()
	key := "TestCache_WriteHash"
	want := map[string]any{
		"value1": "TestCache_ReadHash_OK_Value1",
		"value2": "TestCache_ReadHash_OK_Value2",
	}
	ttl := 1 * time.Minute

	t.Cleanup(func() {
		repo.cache.Client.Del(ctx, key)
	})

	if err := repo.WriteHash(ctx, key, want, ttl); err != nil {
		t.Fatal(err)
	}

	got, err := repo.cache.Client.HGetAll(ctx, key).Result()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("-want +got:\n%s", diff)
	}
}
